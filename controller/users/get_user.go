package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mynameisarwan/test_case/common/models"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user []models.User

	if result := h.DB.Find(&user, "\"Userid\" = ?", id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &user)
}
