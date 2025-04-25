package idempotency

import (
	"bytes"
	"hausparty/libs/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IdempotencyMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Idempotency-Key")
		if key == "" {
			c.Next() // proceed normally
			return
		}

		var existing models.IdempotencyKey
		if err := db.Where("key = ?", key).First(&existing).Error; err == nil {
			// Key found – return cached response
			c.Data(existing.ResponseStatus, "application/json", []byte(existing.ResponseBody))
			c.Abort()
			return
		}

		// Buffer the response for saving later
		writer := &responseBuffer{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
		c.Writer = writer
		c.Next()

		// After request, store response
		db.Create(&models.IdempotencyKey{
			Key:            key,
			ResponseBody:   writer.body.String(),
			ResponseStatus: writer.status,
		})
	}
}

type responseBuffer struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w *responseBuffer) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseBuffer) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
