package run

import (
	"github.com/NethermindEth/eigen-wiz/internal/commands"
	"github.com/NethermindEth/eigen-wiz/internal/compose"
	"github.com/NethermindEth/eigen-wiz/internal/data"
)

// Run starts the package at the given path.
func Run(dataDir *data.DataDir, instance string) error {
	i, err := dataDir.Instance(instance)
	if err != nil {
		return err
	}
	err = i.Lock()
	if err != nil {
		return err
	}
	composePath := i.ComposePath()

	cmdRunner := commands.NewCMDRunner()
	dockerCompose := compose.NewComposeManager(&cmdRunner)
	return dockerCompose.Up(compose.DockerComposeUpOptions{
		Path: composePath,
	})
}