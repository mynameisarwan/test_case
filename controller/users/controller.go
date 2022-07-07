package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/MyWeb")
	routes.GET("/DisplayAllUsers", h.GetUsers)
	routes.GET("/DisplayUser/:id", h.GetUser)
}
