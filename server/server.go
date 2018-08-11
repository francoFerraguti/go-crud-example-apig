package server

import (
	"github.com/francoFerraguti/go-crud-example-apig/middleware"
	"github.com/francoFerraguti/go-crud-example-apig/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	router.Initialize(r)
	return r
}
