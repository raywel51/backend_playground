package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var responseBuffer bytes.Buffer

		// Use the custom response writer for the context
		c.Writer = &CustomResponseWriter{c.Writer, &responseBuffer}
		// Process the request
		c.Next()

		// Log the request information
		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		requestParams := c.Request.URL.Query()
		requestBody := c.Request.PostForm
		responseBody := responseBuffer.String()

		logEntry := fmt.Sprintf("[GIN] %s | %3d | %s | %s | %s %s", end.Format("2006/01/02 - 15:04:05"), status, latency, clientIP, method, path)
		if len(requestParams) > 0 {
			logEntry += fmt.Sprintf("\n"+" | Params: %v", requestParams)
		}
		if len(requestBody) > 0 {
			logEntry += fmt.Sprintf("\n"+" | Body: %v", requestBody)
		}
		logEntry += fmt.Sprintf("\n"+" | Res: %v", responseBody)

		fmt.Println("\n============================\n" + logEntry + "\n")
	}
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	BodyWriter io.Writer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	if err == nil {
		_, err = w.BodyWriter.Write(b)
	}
	return n, err
}
