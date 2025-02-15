package agent

import (
	"errors"

	"github.com/becker63/simtricky/common/logger"
	"github.com/becker63/simtricky/common/pb"
)

var (
	ErrInvalidMachineID = errors.New("invalid machine ID")
)

func StartDockerServer(req *pb.Config) error {
	if req.MachineId == 1 {
		log := logger.GetLogger()
		log.Debug(req.Dockerfile)
		return ErrInvalidMachineID
	}
	return nil
}
