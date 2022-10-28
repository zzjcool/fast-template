package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zzjcool/fast-template/render"
)

var (
	valueFilename    = "example/value.yaml"
	templateFilename = "example/README.template"
	outputFilename   = "example/README.md"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build a template",
	Run: func(cmd *cobra.Command, args []string) {
		render.BuildTemplate[render.Config](valueFilename, templateFilename, outputFilename)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.PersistentFlags().StringVarP(&valueFilename, "value", "v", "example/value.yaml", "value file")
	buildCmd.PersistentFlags().StringVarP(&templateFilename, "template", "t", "example/README.template", "template file")
	buildCmd.PersistentFlags().StringVarP(&outputFilename, "output", "o", "example/README.md", "output file")

}
