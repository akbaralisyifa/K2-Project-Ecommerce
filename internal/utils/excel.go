package utils

import (
	"bytes"
	"ecommerce/internal/features/products"
	"strconv"

	"github.com/xuri/excelize/v2"
)


type ExcelUtilityInterface interface{
	DownloadExcel(data []products.Product) ([]byte, error)
}

type excelUntility struct {}

func NewExcelUtility() ExcelUtilityInterface{
	return &excelUntility{}
}

func (xu *excelUntility) DownloadExcel(data []products.Product) ([]byte, error) {
	f := excelize.NewFile();

	sheet := "Products";
	index, _ := f.NewSheet(sheet);

	// hapus sheet pertama yang di buat secara otomatis 
	f.DeleteSheet("Sheet1")

	// set sheet nya menjadi yang pertama
	f.SetActiveSheet(index)

	// header 
	headers := []string{"id", "UserID", "Name", "Category", "Description", "Price", "Stock", "ImageUrl"};

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1);
		f.SetCellValue(sheet, cell, header)
	};

	for i, product := range data{
		f.SetCellValue(sheet, "A"+strconv.Itoa(i+2), product.ID)
		f.SetCellValue(sheet, "B"+strconv.Itoa(i+2), product.UserID)
		f.SetCellValue(sheet, "C"+strconv.Itoa(i+2), product.Name)
		f.SetCellValue(sheet, "D"+strconv.Itoa(i+2), product.Category)
		f.SetCellValue(sheet, "E"+strconv.Itoa(i+2), product.Description)
		f.SetCellValue(sheet, "F"+strconv.Itoa(i+2), product.Price)
		f.SetCellValue(sheet, "G"+strconv.Itoa(i+2), product.Stock)
		f.SetCellValue(sheet, "H"+strconv.Itoa(i+2), product.ImageUrl)
	};

	var buf bytes.Buffer;

    err := f.Write(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}