package cmd

import (
	"errors"
	"fmt"
	"fullstack/cmd/version"
	"fullstack/common/global"
	"fullstack/pkg"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "fullstack",
	Short:        "fullstack",
	SilenceUsage: true,
	Long:         `fullstack`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageMsg := `欢迎使用 ` + pkg.Green(`fullstack `+global.Version) + ` 版本`
	usageUrl := `欢迎访问 ` + pkg.Green(global.Url) + ` ` + global.Name
	fmt.Printf("%s\n", usageMsg)
	fmt.Printf("%s\n", usageUrl)
}

func init() {
	rootCmd.AddCommand(version.StartCmd)
}

//Execute : apply commands
func Exec() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
