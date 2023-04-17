package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(param gin.LogFormatterParams) string {

	return fmt.Sprintf("\n IP: %s,  TIME: [%s] - METHOD: %s | PATH: %s, LATENCY: %v", param.ClientIP, param.TimeStamp.Format(time.RFC822), param.Method, param.Path, param.Latency)
}
