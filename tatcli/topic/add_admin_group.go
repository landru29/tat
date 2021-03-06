package topic

import (
	"github.com/ovh/tat/tatcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdTopicAddAdminGroup.Flags().BoolVarP(&recursive, "recursive", "r", false, "Apply Rights Admin recursively")
}

var cmdTopicAddAdminGroup = &cobra.Command{
	Use:   "addAdminGroup",
	Short: "Add Admin Groups to a topic: tatcli topic addAdminGroup [--recursive] <topic> <groupname1> [groupname2]...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			err := internal.Client().TopicAddAdminGroups(args[0], args[1:], recursive)
			internal.Check(err)
		} else {
			internal.Exit("Invalid argument: tatcli topic addAdminGroup --help\n")
		}
	},
}
