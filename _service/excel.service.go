package _service

import (
	"encoding/csv"
	"net/http"

	"bmwgroup.net/qt56157/excel2json/_handler"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type ExcelData struct {
	Sheet string     `json:"sheet"`
	Data  [][]string `json:"data"`
}

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "service up")
}

func ParseExcelUpload(ctx *gin.Context) {

	file, _, err := ctx.Request.FormFile("upload")
	var xd ExcelData
	var xdlist []ExcelData
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	excelFile, err := excelize.OpenReader(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	sheets := excelFile.GetSheetList()

	for i := 0; i < len(sheets); i++ {
		rows := _handler.GetRowsFromSheet(excelFile, sheets[i])
		xd.Sheet = sheets[i]
		xd.Data = rows
		xdlist = append(xdlist, xd)
	}

	ctx.JSON(http.StatusOK, xdlist)

}

type CSVData struct {
	Header []string                 `json:"header"`
	Rows   []map[string]interface{} `json:"data"`
}

func ParseCSVUpload(ctx *gin.Context) {
	var csvData CSVData
	file, _, err := ctx.Request.FormFile("upload")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	csvData.Header = records[0]

	for x := 1; x < len(records); x++ {
		row := make(map[string]interface{})
		for y := 0; y < len(records[0]); y++ {
			key := records[0][y]
			value := records[x][y]
			row[key] = value
		}
		csvData.Rows = append(csvData.Rows, row)
	}
	ctx.JSON(http.StatusOK, csvData)

}
