package testing

import (
	"testing"

	dockercontainer "github.com/moby/moby/api/types/container"
)

type IsReadyFunc func(Instance) bool

type TestFunc func(*testing.T, Instance)

type Version struct {
	Image string
	ENV   []string
	Cmd   []string
}

func ParallelTest(t *testing.T, versions []Version, readyFn IsReadyFunc, testFn TestFunc) {
	_ = "STUB: not implemented"
	return
}

func containerLogs(t *testing.T, c *DockerContainer) []byte { _ = "STUB: not implemented"; return nil }

type Instance interface {
	Host() string
	Port() uint
	PortFor(int) uint
	NetworkSettings() dockercontainer.NetworkSettings
	KeepForDebugging()
}
