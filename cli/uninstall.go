package cli

import (
	"github.com/NethermindEth/egn/pkg/daemon"
	"github.com/spf13/cobra"
)

func UninstallCmd(d daemon.Daemon) *cobra.Command {
	var instanceId string
	cmd := cobra.Command{
		Use:   "uninstall [INSTANCE_ID]",
		Short: "Uninstall an instance",
		Long:  "Uninstall an instance. This will stop the instance and remove all its data. INSTANCE_ID is required as the unique argument, and it is the combination of the instance repository name and the instance tag computed during the installation, like this: [repository-name]-[tag].",
		Args:  cobra.ExactArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			instanceId = args[0]
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return d.Uninstall(instanceId)
		},
	}
	return &cmd
}