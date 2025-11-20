package handlers

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shadykip/go-worker/internal/models"
	"gorm.io/gorm"
)

type EnqueueRequest struct {
	Type    string          `json:"type" binding:"required"`
	Payload json.RawMessage `json:"payload" binding:"required"`
}

func EnqueueJob(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req EnqueueRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		job := models.Job{
			Type:    req.Type,
			Payload: string(req.Payload), // Store as string for JSONB
			Status:  "pending",
		}

		if err := db.Create(&job).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to enqueue job"})
			return
		}

		c.JSON(202, gin.H{
			"job_id":    job.ID,
			"type":      job.Type,
			"status":    job.Status,
			"queued_at": job.CreatedAt.Format(time.RFC3339),
		})
	}
}
