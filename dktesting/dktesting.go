package dktesting

import (
	"testing"

	"github.com/dhui/dktest"
)

type ContainerSpec struct {
	ImageName string
	Options   dktest.Options
}

func (s *ContainerSpec) Cleanup() (retErr error) { _ = "STUB: not implemented"; return nil }

func ParallelTest(t *testing.T, specs []ContainerSpec,
	testFunc func(*testing.T, dktest.ContainerInfo)) {
	_ = "STUB: not implemented"
	return
}
