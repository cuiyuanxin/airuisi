package middleware

import (
	"fmt"
	"time"

	"github.com/cuiyuanxin/airuisi/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		beginTime := time.Now().Unix()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now().Unix()
		// 执行时间
		latencyTime := endTime - beginTime

		global.Logger.Info(
			fmt.Sprintf("access log: method: %s, status_code: %d, request_url: %s, client_ip: %s, begin_time: %d, end_time: %d, latency_time: %d",
				c.Request.Method,     // 请求方式
				c.Writer.Status(),    // 状态码
				c.Request.RequestURI, // 请求路由
				c.ClientIP(),         // 请求ip
				beginTime,
				endTime,
				latencyTime,
			),
			zap.String("trace_id", c.MustGet("X-Trace-ID").(string)),
			zap.String("span_id", c.MustGet("X-Span-ID").(string)),
		)
	}
}
