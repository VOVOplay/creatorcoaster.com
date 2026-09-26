package database

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadStaticMarkdownFile(fileName string) string {
	filePath := filepath.Join(fmt.Sprintf("./config/static-markdown-content/%v", fileName))
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	data_str := string(data)

	return data_str
}
