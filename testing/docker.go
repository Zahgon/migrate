package testing

import (
	"io"
	"testing"

	dockercontainer "github.com/moby/moby/api/types/container"
	dockerclient "github.com/moby/moby/client"
)

func NewDockerContainer(t testing.TB, image string, env []string, cmd []string) (*DockerContainer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DockerContainer struct {
	t                  testing.TB
	client             *dockerclient.Client
	ImageName          string
	ENV                []string
	Cmd                []string
	ContainerId        string
	ContainerName      string
	ContainerJSON      dockercontainer.InspectResponse
	containerInspected bool
	keepForDebugging   bool
}

func (d *DockerContainer) PullImage() (err error) { _ = "STUB: not implemented"; return nil }

func (d *DockerContainer) Start() error { _ = "STUB: not implemented"; return nil }

func (d *DockerContainer) KeepForDebugging() { _ = "STUB: not implemented"; return }

func (d *DockerContainer) Remove() error { _ = "STUB: not implemented"; return nil }

func (d *DockerContainer) Inspect() error { _ = "STUB: not implemented"; return nil }

func (d *DockerContainer) Logs() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (d *DockerContainer) portMapping(selectFirst bool, cPort int) (hostIP string, hostPort uint, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (d *DockerContainer) Host() string { _ = "STUB: not implemented"; return "" }

func (d *DockerContainer) Port() uint { _ = "STUB: not implemented"; return 0 }

func (d *DockerContainer) PortFor(cPort int) uint { _ = "STUB: not implemented"; return 0 }

func (d *DockerContainer) NetworkSettings() dockercontainer.NetworkSettings {
	_ = "STUB: not implemented"
	return *new(dockercontainer.NetworkSettings)
}

type dockerImagePullOutput struct {
	Status          string `json:"status"`
	ProgressDetails struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
	Id       string `json:"id"`
	Progress string `json:"progress"`
}

func pseudoRandStr(n int) string { _ = "STUB: not implemented"; return "" }
