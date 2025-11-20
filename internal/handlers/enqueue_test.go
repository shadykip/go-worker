package handlers

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Use in-memory SQLite for unit tests (or clean PostgreSQL)
	// For simplicity, reuse worker DB with cleanup
	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	assert.NoError(t, err)
	db.Exec("DELETE FROM jobs")
	return db
}

func TestEnqueueJob_Valid(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB(t)
	defer db.Exec("DELETE FROM jobs")

	r := gin.New()
	r.POST("/jobs", EnqueueJob(db))

	jsonBody := `{
		"type": "test_job",
		"payload": {"key": "value"}
	}`
	req := httptest.NewRequest("POST", "/jobs", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 202, resp.Code)
	assert.Contains(t, resp.Body.String(), "job_id")
}
