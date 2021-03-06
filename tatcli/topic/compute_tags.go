package topic

import (
	"github.com/ovh/tat"
	"github.com/ovh/tat/tatcli/internal"
	"github.com/spf13/cobra"
)

var cmdTopicComputeTags = &cobra.Command{
	Use:   "computetags",
	Short: "Compute Tags on this topic, only for tat admin and administrators on topic : tatcli topic computetags <topic>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			out, err := internal.Client().TopicComputeTags(tat.TopicNameJSON{Topic: args[0]})
			internal.Check(err)
			if internal.Verbose {
				internal.Print(out)
			}
		} else {
			internal.Exit("Invalid argument: tatcli topic computetags --help\n")
		}
	},
}
