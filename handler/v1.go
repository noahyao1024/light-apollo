package handler

import (
	"encoding/json"
	"fmt"
	"light-apollo/storage"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

const lruTTL = 60 * time.Second

var lru *cache.Cache

func init() {
	lru = cache.New(lruTTL, 10*time.Minute)
}

type Release struct {
	AppID          string `gorm:"primaryKey;Column:AppId" json:"appId"`
	Cluster        string `gorm:"Column:ClusterName"`
	Namespace      string `gorm:"Column:NamespaceName"`
	ReleaseKey     string `gorm:"Column:ReleaseKey"`
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

	// Get from cache.
	key := fmt.Sprintf("%s:%s:%s", appID, cluster, namespace)
	if raw, exists := lru.Get(key); exists {
		if data, ok := raw.(map[string]interface{}); ok {
			c.JSON(200, data)
			return
		}
	}

	releases := make([]Release, 0)
	err := storage.GetDB().Where(&Release{
		AppID:     appID,
		Cluster:   cluster,
		Namespace: namespace,
	}).Find(&releases).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, release := range releases {
		if release.AppID == appID && release.Cluster == cluster && release.Namespace == namespace {
			data := make(map[string]interface{})
			json.Unmarshal([]byte(release.Configurations), &data)

			// Set to cache.
			lru.Set(key, data, lruTTL)

			c.JSON(200, data)
			return
		}
	}

	c.JSON(200, gin.H{})
}
