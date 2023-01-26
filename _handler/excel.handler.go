package _handler

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func OpenExcel(fileName string) (*excelize.File, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {

		return nil, err
	}
	return f, err
}

func GetRowsFromSheet(excelFile *excelize.File, sheetName string) [][]string {

	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)

	}
	return rows
}

func GetSheetList(excelFile *excelize.File, sheetName string) [][]string {

	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)

	}
	return rows
}
