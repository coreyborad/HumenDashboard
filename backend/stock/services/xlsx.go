package services

import (
	"fmt"
	"stock/models"
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
	fmt.Println(info)
	return nil
}
