package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(context *Context) {
		// Start timer
		t := time.Now()
		// Process request
		context.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", context.StatusCode, context.Req.RequestURI, time.Since(t))
	}
}
