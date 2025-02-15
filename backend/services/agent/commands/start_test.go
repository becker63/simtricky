package agent

import (
	"testing"

	"github.com/becker63/simtricky/common/pb"
	"github.com/stretchr/testify/assert"
)

func TestStartAgentLogic(t *testing.T) {
	t.Run("Valid Machine ID", func(t *testing.T) {
		err := StartDockerServer(&pb.Config{
			MachineId: 1,
		})
		assert.Nil(t, err)
	})

	t.Run("Valid Machine ID", func(t *testing.T) {
		err := StartDockerServer(&pb.Config{
			MachineId: 0,
		})
		assert.Nil(t, err)
	})
}
