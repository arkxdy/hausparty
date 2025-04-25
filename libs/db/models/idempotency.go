package models

import "time"

type IdempotencyKey struct {
	ID             uint   `gorm:"primaryKey"`
	Key            string `gorm:"uniqueIndex"`
	ResponseBody   string
	ResponseStatus int
	CreatedAt      time.Time
}
