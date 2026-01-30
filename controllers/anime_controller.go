package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-quickstart/config"
	"gin-quickstart/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET ALL dengan Cache
func GetAllAnime(c *gin.Context) {
	ctx := context.Background()
	cacheKey := "anime:all"

	// 1. Cek apakah data ada di Redis
	cachedData, err := config.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data ada di cache, langsung return
		var anime []models.Anime
		json.Unmarshal([]byte(cachedData), &anime)
		c.JSON(http.StatusOK, gin.H{
			"source": "cache",
			"data":   anime,
		})
		return
	}

	// 2. Data tidak ada di cache, ambil dari database
	var anime []models.Anime
	config.DB.Find(&anime)

	// 3. Simpan ke Redis cache (expired 5 menit)
	animeJSON, _ := json.Marshal(anime)
	config.RedisClient.Set(ctx, cacheKey, animeJSON, 5*time.Minute)

	c.JSON(http.StatusOK, gin.H{
		"source": "database",
		"data":   anime,
	})
}

// GET BY ID dengan Cache
func GetByIdAnime(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	cacheKey := fmt.Sprintf("anime:%s", id)

	// Cek cache
	cachedData, err := config.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var anime models.Anime
		json.Unmarshal([]byte(cachedData), &anime)
		c.JSON(http.StatusOK, gin.H{
			"source": "cache",
			"data":   anime,
		})
		return
	}

	// Ambil dari database
	var anime models.Anime
	if err := config.DB.First(&anime, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "tidak ditemukan"})
		return
	}

	// Simpan ke cache
	animeJSON, _ := json.Marshal(anime)
	config.RedisClient.Set(ctx, cacheKey, animeJSON, 5*time.Minute)

	c.JSON(http.StatusOK, gin.H{
		"source": "database",
		"data":   anime,
	})
}

// CREATE - Hapus cache setelah create
func CreateAnime(c *gin.Context) {
	var anime models.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&anime)

	// Hapus cache karena data baru ditambahkan
	ctx := context.Background()
	config.RedisClient.Del(ctx, "anime:all")

	c.JSON(http.StatusOK, anime)
}

// UPDATE - Hapus cache setelah update
func UpdateAnime(c *gin.Context) {
	var anime models.Anime
	id := c.Param("id")

	if err := config.DB.First(&anime, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "tidak ditemukan"})
		return
	}

	c.ShouldBindJSON(&anime)
	config.DB.Save(&anime)

	// Hapus cache
	ctx := context.Background()
	config.RedisClient.Del(ctx, "anime:all", fmt.Sprintf("anime:%s", id))

	c.JSON(http.StatusOK, anime)
}

// DELETE - Hapus cache setelah delete
func DeleteAnime(c *gin.Context) {
	var anime models.Anime
	id := c.Param("id")

	if err := config.DB.First(&anime, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "tidak ditemukan"})
		return
	}

	config.DB.Delete(&anime)

	// Hapus cache
	ctx := context.Background()
	config.RedisClient.Del(ctx, "anime:all", fmt.Sprintf("anime:%s", id))

	c.JSON(http.StatusOK, gin.H{"message": "anime berhasil dihapus"})
}
