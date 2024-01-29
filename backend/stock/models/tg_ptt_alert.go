package models

type TgPttAlert struct {
	ID      uint64 `json:"id" gorm:"primaryKey"`
	ChartID string `json:"chart_id"`
	Kanban  string `json:"kanban"`
	KeyWord string `json:"key_word"`
}

// TableName overrides the table name used by Device to `hmi`
func (TgPttAlert) TableName() string {
	return "tg_ptt_alert"
}
