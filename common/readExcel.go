package common

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
)

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func GetFileByFilename(filename string, ifCreate bool) (*excelize.File, error) {
	var excel *excelize.File
	var err error
	if FileExists(filename) {
		excel, err = excelize.OpenFile(filename)
		if err != nil {
			return nil, err
		}
	}
	return excel, nil
}

//根据指定的表明读取数据
func ReadExcelAndPrintall(filename string, sheetname string) {
	excel, err := GetFileByFilename(filename, false)

	if err != nil {
		fmt.Println(err)
	}
	var j int32
	for i := 1; ; i++ {
		for j = 'A'; ; j++ {
			value := excel.GetCellValue(sheetname, string(j)+strconv.Itoa(i))
			if value == "" {
				if j == 'A' {
					return
				} else {
					break
				}
			}
			fmt.Printf("%8s\t", value)
		}
		fmt.Println()
	}
}
