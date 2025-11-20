package models

import "time"

type Job struct {
	ID          uint   `gorm:"primarykey"`
	Type        string `gorm:"not null"`                   // e.g., "send_email"
	Payload     string `gorm:"type:jsonb"`                 // JSON payload
	Status      string `gorm:"not null;default:'pending'"` // pending, processing, success, failed
	Retries     int    `gorm:"default:0"`
	FailedAt    *time.Time
	CreatedAt   time.Time
	ProcessedAt *time.Time
}
