package controller

import (
	stderrors "errors"
	"net/http"

	bizerrors "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"github.com/gin-gonic/gin"
)

// respondError 把业务异常按约定结构返回，未知异常统一按 500 处理。
func respondError(c *gin.Context, err error) {
	var be bizerrors.BusinessError
	if stderrors.As(err, &be) {
		status := be.Status
		if status == 0 {
			status = http.StatusBadRequest
		}
		c.JSON(status, be)
		return
	}
	logger.Error("unexpected error:", err)
	c.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_ERROR", "message": "服务器内部错误"})
}
