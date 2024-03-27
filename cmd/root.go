package cmd

import (
	"fmt"
	"os"

	"github.com/danielmesquitta/webpfyer/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "txtfyer",
	Short: "Converts png, jpg and jpeg files to webp.",
	Long:  `Transforms file folders or single files into webp.`,

	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			fmt.Println(err)
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
		quality, err := cmd.Flags().GetUint("quality")
		if err != nil {
			fmt.Println(err)
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
		err = usecase.ConvertToWebp(path, quality)
		if err != nil {
			fmt.Println("Invalid command")
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("path", "p", "./", "Path to a file folder or a single file to convert to webp.")
	rootCmd.Flags().UintP("quality", "q", 80, "Quality specify the compression factor for RGB channels between 0 and 100.")
}
