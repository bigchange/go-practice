package controller

import (
	"net/http"
	"strconv"

	"github.com/bigchange/go-practice/testcase/db"
	"github.com/bigchange/go-practice/testcase/middleware"
	"github.com/gin-gonic/gin"
)
func Handler1 (c *gin.Context) {
	_ = middleware.GetHubFromCtx(c)
	vStr := c.DefaultQuery("k", "0")
  v, err := strconv.ParseInt(vStr, 10, 64)
  if err != nil {
  	c.JSON(http.StatusBadRequest, err)
  	return
	}
	status, err := db.GetDB().GetData(c.Request.Context(), v)
	// status, err := hub.PrimaryDB.GetData(c.Request.Context(), v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status.Status})
	return
}
