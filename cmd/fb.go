package cmd

import (
	"fmt"
	"os"
	"strconv"
	"ztool/internal/fb"

	"github.com/spf13/cobra"
)

var (
	fbStr     string
	tableNum  int
	fbByMod   bool
	fbByCrc32 bool

	formatWidth int
)

var FbCmd = &cobra.Command{
	Use:   "fb",
	Short: "查询分表后缀",
	Run: func(cmd *cobra.Command, args []string) {
		var result int
		if fbByCrc32 {
			result = fb.FbByCrc32(fbStr, tableNum)
			fmt.Println("分表模式:crc3")
		} else {
			fmt.Println("分表模式:直接取模")
			result = fb.FbByMod(fbStr, tableNum)
		}

		var resutlStr string
		if formatWidth == 0 {
			resutlStr = fmt.Sprintf("%d", result)
		} else {
			format := fmt.Sprintf("%%0%dd", formatWidth)
			resutlStr = fmt.Sprintf(format, result)
		}

		fmt.Printf("分表数量:%d, 分表:%s\n", tableNum, resutlStr)
	},
}

const (
	defaultTableNumEnv    = "ztool_tablenumber"
	defaultFormatWidthEnv = "ztool_formatwidth"
)

func init() {
	defaultTableNum, _ := strconv.Atoi(os.Getenv(defaultTableNumEnv))
	defaultFormatWidth, _ := strconv.Atoi(os.Getenv(defaultFormatWidthEnv))

	if defaultTableNum == 0 {
		defaultTableNum = 1
	}

	FbCmd.Flags().StringVarP(&fbStr, "str", "s", "", "分表字段")
	FbCmd.Flags().IntVarP(&tableNum, "tableNum", "n", defaultTableNum, "分表数,可通过$"+defaultTableNumEnv+"设置默认值")
	FbCmd.Flags().BoolVarP(&fbByMod, "fbByMod", "m", false, "直接取模(默认模式)")
	FbCmd.Flags().BoolVarP(&fbByCrc32, "fbByCrc32", "c", false, "CRC32哈希后取模")
	FbCmd.Flags().IntVarP(&formatWidth, "width", "w", defaultFormatWidth,
		"输出结果宽度,0表示不设置,可通过$"+defaultFormatWidthEnv+"设置默认值")
}
