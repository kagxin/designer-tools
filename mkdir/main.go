package main

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ExcelName 文件夹名字
var (
	SheetName = "Sheet1"
	FilePath  = "mkdir.xlsx"
)

// GetNameFromExcel c
func GetNameFromExcel(filePath, sheetName string) ([]string, error) {
	texts := make([]string, 0, 2048)
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows := f.GetRows(sheetName)
	for _, row := range rows {
		texts = append(texts, row...)
	}
	return texts, nil
}

func main() {
	var ts string
	dirNames, err := GetNameFromExcel(FilePath, SheetName)
	if err != nil {
		fmt.Println("未找到文件mkdir.xlsx。")
		fmt.Scanf("%s", &ts)
		os.Exit(-1)

	}
	if len(dirNames) == 0 {
		fmt.Println("请检查Sheet1是否为空。")
		fmt.Scanf("%s", &ts)
		os.Exit(-1)
	}
	for _, dir := range dirNames {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm)
		}
	}
	fmt.Printf("完成，共%d个文件夹\n", len(dirNames))
	fmt.Scanf("%s", &ts)
}
