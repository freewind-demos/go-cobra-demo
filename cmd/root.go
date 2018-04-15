package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// `go-cobra-demo`是这个命令行工具的名字（而不是里面的一个command），它应该跟对应的目录同名（但本项目使用了不同的名字`go-cobra-demo`）
//  下面的介绍也是针对这个工具本身的
// 可以通过 `hugo --help` 查看到
var rootCmd = &cobra.Command{
	Use:   "go-cobra-demo",
	Short: "Simple demo of cobra",
	Long:  `This is a very simple demo for cobra, you can see more at: https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello cobra!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
