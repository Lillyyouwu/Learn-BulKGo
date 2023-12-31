package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	bulkdo "github.com/Lillyyouwu/Learn-BulkGo"
)

// BulkDoHome is the home of bulkdo
const BulkDoHome = ".bulkdo"

// DefaultItemFile is default params
const DefaultItemFile = ".bulkdoitems"

func main() {
	// 检查长度
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	fmt.Printf("Execute command '%v'", cmdName)

	// 获取用户主目录
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 命令文件的目录 用字符串连接起来
	cmdTplPath := filepath.Join(user.HomeDir, BulkDoHome, cmdName+".tpl")
	if _, err := os.Stat(cmdTplPath); os.IsNotExist(err) {
		fmt.Printf("Command Template path '%v' does not exist.", cmdTplPath)
		os.Exit(1)
	}

	if _, err := os.Stat(DefaultItemFile); os.IsNotExist(err) {
		fmt.Printf("DefaultItemFile path '%v' does not exist.", DefaultItemFile)
		os.Exit(1)
	}

	cmdTplFile, cmdTplFileErr := os.Open(cmdTplPath)
	if cmdTplFileErr != nil {
		fmt.Println(cmdTplFileErr)
		os.Exit(1)
	}

	itemFile, itemFileErr := os.Open(DefaultItemFile)
	if itemFileErr != nil {
		fmt.Println(itemFileErr)
		os.Exit(1)
	}

	outs, outsErr := bulkdo.BulkDo(cmdTplFile, itemFile)
	if outsErr != nil {
		fmt.Println(outsErr)
		os.Exit(1)
	}

	for _, out := range outs {
		fmt.Println(out)
	}

}
