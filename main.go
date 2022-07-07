package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mynameisarwan/test_case/common/db"
	"github.com/mynameisarwan/test_case/controller/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"port":  port,
	// 		"dbUrl": dbUrl,
	// 	})
	// })
	users.RegisterRoutes(r, h)

	r.Run(port)
}
