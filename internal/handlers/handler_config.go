package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegistryRoutes(r *gin.Engine, db *gorm.DB) {
	
	registryProjectRoutes(r, db)

}
