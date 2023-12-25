package handler

import (
	"light-apollo/storage"

	"github.com/gin-gonic/gin"
)

type Release struct {
	AppID          string `gorm:"primaryKey" json:"appId"`
	Cluster        string `gorm:"primaryKey" json:"cluster"`
	Namespace      string `gorm:"primaryKey" json:"namespace"`
	ReleaseKey     string `json:"releaseKey"`
	Configurations string `gorm:"type:text" json:"configurations"`
}

func (Release) TableName() string {
	return "Release"
}

// Configs mock as {{apollo_host}}/configfiles/json/your_app/default/application
func Configs(c *gin.Context) {
	appID := c.Param("app_id")
	cluster := c.Param("cluster")
	namespace := c.Param("namespace")

	releases := make([]Release, 0)
	err := storage.GetDB().Where(&Release{}).Find(&releases).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"app_id":    appID,
		"cluster":   cluster,
		"namespace": namespace,
		"releases":  releases,
	})
}
