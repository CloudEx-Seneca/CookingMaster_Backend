package user

import (
	"CookingMaster_Backend/app/usercenter/api/internal/svc"
	"CookingMaster_Backend/app/usercenter/api/internal/types"
	"CookingMaster_Backend/app/usercenter/model"
	"CookingMaster_Backend/app/usercenter/rpc/usercenterClient"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/argon2"
	"net"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userId, err := generateUserId()
	if err != nil {
		return &types.RegisterResp{}, err
	}
	gtResp, err := l.svcCtx.UserCenterRpc.GenerateToken(l.ctx, &usercenterClient.GenerateTokenReq{
		UserId:    userId,
		TokenType: model.RegisterTokenType,
	})
	if err != nil {
		logx.Errorf("generate token error: %v", err)
		return &types.RegisterResp{}, err
	}

	passwordHash, err := GeneratePassword(req.Password)
	if err != nil {
		return &types.RegisterResp{}, err
	}
	l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		Id:       userId,
		Email:    req.Email,
		Password: passwordHash,
	})

	return &types.RegisterResp{
		RegisterToken:  gtResp.Token,
		RegisterExpire: gtResp.Expire,
	}, nil
}

func GeneratePassword(password string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	time := uint32(3)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)
	encoded := fmt.Sprintf("%s$%d$%d$%d$%s$%s",
		"argon2id",
		time,
		memory,
		threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encoded, nil
}

func generateUserId() (int64, error) {
	nodeId, err := generateNodeId()
	if err != nil {
		logx.Errorf("generateNodeId err: %s", err.Error())
		return 2025, err
	}

	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		logx.Errorf("snowflake.NewNode err: %v", err)
		return 2025, err
	}

	return node.Generate().Int64(), nil
}

func generateNodeId() (int64, error) {
	mac, err := getMacAddress()
	if err != nil {
		logx.Errorf("getMacAddress err: %s", err.Error())
		return 1023, err
	}

	hash := sha256.Sum256([]byte(mac))
	nodeId := int64(binary.BigEndian.Uint64(hash[:8]))
	if nodeId < 0 {
		nodeId = -nodeId
	}

	nodeId = nodeId % 1024
	return nodeId, nil
}

func getMacAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback == 0 && len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String(), nil
		}
	}

	return "", fmt.Errorf("can not find interface by macAddress")
}
