package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mynameisarwan/test_case/common/models"
)

func (h handler) GetUsers(c *gin.Context) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &users)
}
