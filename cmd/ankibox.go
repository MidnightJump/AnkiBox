/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	utils2 "AnkiBox/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var mdFilePath string
var deckName string

// ankiboxCmd represents the ankibox command
var ankiboxCmd = &cobra.Command{
	Use:   "ankibox",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addNotes(mdFilePath, deckName)
	},
}

func addNotes(filePath string, deckName string) {
	sections, err := utils2.ParseMarkdown(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	content := utils2.GetAnkiContentFromSection(sections)
	err = utils2.AddNote(deckName, content)
	if err != nil {
		fmt.Printf("%e", err)
	}
}
func init() {
	ankiboxCmd.Flags().StringVarP(&mdFilePath, "file", "f", "./", "请输入要添加的MD文件绝对路径")
	ankiboxCmd.Flags().StringVarP(&deckName, "deck", "d", "base", "请输入要添加的牌组名")
	rootCmd.AddCommand(ankiboxCmd)

}
