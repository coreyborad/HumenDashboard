package services

import (
	"context"
	"errors"
	"fmt"
	"stock/google"
	"stock/models"

	"google.golang.org/api/sheets/v4"
)

// XlsxService Xlsx Service
type XlsxService struct {
}

// NewXlsxService New Xlsx Service
func NewXlsxService() *XlsxService {
	return &XlsxService{}
}

// AppendRecord AppendRecord
func (s *XlsxService) AppendRecord(info models.XlsxForm) error {
	ctx := context.Background()
	spreadsheetList := map[int]string{}
	spreadsheetList[2021] = "1MAyVqhbPytK810aSUun2MSKBMz1u3ARgzBJIe11h0rE"
	spreadsheetList[2022] = "1Itzco7QP0M7V1T0eh_AHfgbHbmcOfnQyMmiEKGEo8SU"

	if spreadsheetList[info.Date.Year()] == "" {
		return errors.New("Not found years")
	}
	spreadsheetId := spreadsheetList[info.Date.Year()]
	googleServ := google.GetService()
	rangeStr := fmt.Sprintf("%s.", info.Date.Month().String()[:3])

	appendValue := []interface{}{
		info.Date.Format("1/2"),
		info.Payer,
		info.Category,
		info.SubCategory,
		info.Note,
		info.Cost,
	}
	val := [][]interface{}{}
	val = append(val, appendValue)
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         val,
	}
	valueInputOption := "USER_ENTERED"
	insertDataOption := "INSERT_ROWS"
	_, err := googleServ.Spreadsheets.Values.Append(spreadsheetId, rangeStr, rb).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil {
		return err
	}
	return nil
}
