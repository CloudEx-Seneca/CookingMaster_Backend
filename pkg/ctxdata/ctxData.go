package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserId = "jwtUserId"
var CtxKeyJwtTokenType = "jwtTokenType"

func GetUidFromCtx(ctx context.Context) int64 {
	return getTokenInfoFromCtx(ctx, CtxKeyJwtUserId)
}

func GetTokenTypeFromCtx(ctx context.Context) int64 {
	return getTokenInfoFromCtx(ctx, CtxKeyJwtTokenType)
}

func getTokenInfoFromCtx(ctx context.Context, key string) int64 {
	var tokenInfo int64
	if jsonTokenInfo, ok := ctx.Value(key).(json.Number); ok {
		if int64tokeninfo, err := jsonTokenInfo.Int64(); err == nil {
			tokenInfo = int64tokeninfo
		} else {
			logx.WithContext(ctx).Error("getTokenInfoFromCtx err : %+v", err)
		}
	} else {
		logx.WithContext(ctx).Error("ctx.Value ok=false ctx: %+v", ctx)
	}

	return tokenInfo
}
