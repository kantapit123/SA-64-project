package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kantapit123/sa-64-ex/entity"
)

// POST /screening_records
func CreateScreeningRecord(c *gin.Context) {

	var screening_record entity.Screening_record
	var patient entity.Patient
	var medical_product entity.Medical_product
	var user entity.User

	//10:ผลลัพธ์ที่ได้จากขั้นตอนที่ x จะถูก bind เข่้าตัวแปร scr
	if err := c.ShouldBindJSON(&screening_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//31:ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", screening_record.Scr_userID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	//7:ค้นหา patient ด้วย p_id
	if tx := entity.DB().Where("id = ?", screening_record.Scr_patientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
		return
	}
	//11:ค้นหา medical_product ด้วย m_id
	if tx := entity.DB().Where("id = ?", screening_record.Scr_medical_productID).First(&medical_product); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Medical Product not found"})
		return
	}
	//12:สร้าง Screening_records(p_id, m_id, u_id, illnesses, detail)
	scr := entity.Screening_record{
		//โยงความสัมพันธ์กับ Entity Patient
		//โยงความสัมพันธ์กับ Entity Medical_product
		//โยงความสัมพันธ์กับ Entity User
		Scr_patient:         patient,
		Scr_medical_product: medical_product,
		Scr_user:            user,
		Illnesses:           screening_record.Illnesses,
		Detail:              screening_record.Detail,
		Queue:               screening_record.Queue,
	}
	//13:บันทึก()
	if err := entity.DB().Create(&scr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scr})
}

//GET /Screening_record/:id
func GetScreening_record(c *gin.Context) {
	var screening_record entity.Screening_record
	id := c.Param("id")
	if err := entity.DB().Preload("Patient").Preload("Medical_product").Preload("User").Raw("SELECT * FROM screening_record WHERE id = ?", id).Find(&screening_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": screening_record})
}

//GET /Screening_records
func ListScreening_record(c *gin.Context) {
	var screening_records []entity.Screening_record
	if err := entity.DB().Preload("Patient").Preload("Medical_product").Preload("User").Raw("SELECT * FROM screening_record").Find(&screening_records).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": screening_records})
}

//DELETE /Screening_record/:id
func DeleteScreening(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM screening_record WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Screening not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

//PATCH /Screening_records
func UpdateScreening(c *gin.Context) {
	var screening_record entity.Screening_record
	if err := c.ShouldBindJSON(&screening_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", screening_record.ID).First(&screening_record); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Screening-Record not found"})
		return
	}

	if err := entity.DB().Save(&screening_record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": screening_record})
}
