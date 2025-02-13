package authhelper

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"net"
)

func GenerateUserId() (int64, error) {
	nodeId, err := generateNodeId()
	if err != nil {
		return 2025, err
	}

	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		return 2025, err
	}

	return node.Generate().Int64(), nil
}

func generateNodeId() (int64, error) {
	mac, err := getMacAddress()
	if err != nil {
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
