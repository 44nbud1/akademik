package pkgsnowflake

import (
	"net"
	"os"
	"strconv"
	"strings"

	sf "github.com/bwmarrin/snowflake"
)

type Snowflake interface {
	GenerateInt64() int64
	GenerateString() string
	Generate() ID
}

type ID struct {
	sf.ID
}

type snowflake struct {
	node *sf.Node
}

func New() (*snowflake, error) {
	nodeID, err := sf.NewNode(getNodeIDFromMachineIP())
	if err != nil {
		return nil, err
	}

	return &snowflake{
		node: nodeID,
	}, nil
}

func (s *snowflake) GenerateInt64() int64 {
	return s.node.Generate().Int64()
}

func (s *snowflake) GenerateString() string {
	return s.node.Generate().String()
}

func (s *snowflake) Generate() ID {
	return ID{
		ID: s.node.Generate(),
	}
}

func getMachineIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String()

			}
		}
	}
	return ""
}

func getNodeIDFromMachineIP() int64 {
	ip := getMachineIP()
	s := strings.Split(ip, ".")

	var slice int64 = 0
	for _, digit := range s {
		i, err := strconv.Atoi(digit)
		if err != nil {
			slice += 0
			continue
		}
		slice += int64(i)
	}

	return slice
}
