package handler

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	rep := gin.H{
		"message": "ok",
		"code":    200,
	}
	c.JSON(200, rep)
}

func GetIndexData(c *gin.Context) {
	rep := gin.H{
		"message": "ok",
		"code":    200,
	}
	c.JSON(200, rep)
}

func Download(c *gin.Context)  {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	f.Write(c.Writer)
}