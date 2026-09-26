package controller

import (
	stderrors "errors"
	"net/http"
	"strconv"

	bizerr "cyskillswap/internal/errors"
	"cyskillswap/internal/logger"
	"github.com/gin-gonic/gin"
)

// writeError 统一把业务错误映射为 HTTP 响应
func writeError(c *gin.Context, err error) {
	var be bizerr.BusinessError
	if stderrors.As(err, &be) {
		c.JSON(be.Status, be)
		return
	}
	logger.Error("internal error:", err)
	c.JSON(http.StatusInternalServerError, bizerr.ErrInternal)
}

func parseID(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, bizerr.ErrInvalidID)
		return 0, false
	}
	return id, true
}
