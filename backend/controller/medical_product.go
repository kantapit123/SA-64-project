package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kantapit123/sa-64-ex/entity"
)

//POST /medical_products
func CreateMed(c *gin.Context) {
	var medical_product entity.Medical_product
	if err := c.ShouldBindJSON(&medical_product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&medical_product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medical_product})
}

//GET /medical_product/:id
func GetMed(c *gin.Context) {
	var medical_product entity.Medical_product
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medical_products WHERE id = ?", id).Scan(&medical_product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medical_product})

}

//GET /medical_products
func ListMeds(c *gin.Context) {
	var medical_products []entity.Medical_product
	if err := entity.DB().Raw("SELECT * FROM medical_products").Scan(&medical_products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medical_products})
}

//DELETE /medical_products/:id
func DeleteMed(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medical_products WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medical_product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//PATCH /medical_products
func UpdateMed(c *gin.Context) {
	var medical_product entity.Medical_product
	if err := c.ShouldBindJSON(&medical_product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medical_product.ID).First(&medical_product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medical_product not found"})
		return
	}
	if err := entity.DB().Save(&medical_product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medical_product})
}
