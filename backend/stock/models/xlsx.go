package models

import "time"

type XlsxForm struct {
	Date        *time.Time `json:"date" binding:"required"`
	Payer       string     `json:"payer" binding:"required"`
	Category    string     `json:"category" binding:"required"`
	SubCategory string     `json:"subCategory" binding:"required"`
	Note        string     `json:"note,omitempty"`
	Cost        int64      `json:"cost,omitempty"`
}
