package masyarakatcontrollers

import (
	"Backend_TA/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {
	var masyarakat []models.Masyarakat
	if err := models.DB.Preload("User").Find(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{"data": masyarakat})
}

func ShowId(c *fiber.Ctx) error {
	nik := c.Params("nik")

	if nik == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "NIK kosong"})
	}

	tx := models.DB
	var masyarakat models.Masyarakat
	if err := tx.Preload("User").Joins("JOIN User ON masyarakat.nik = user.id").Where("user.id = ?", nik).First(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(&masyarakat)
}

func UpdateProfile(c *fiber.Ctx) error {
	tx := models.DB
	nik := c.Params("nik")

	if nik == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "NIK required"})
	}

	var cekData models.Masyarakat

	if err := tx.Where("nik = ?", nik).First(&cekData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": err.Error()})
	}

	var masyarakat models.Masyarakat
	if err := c.BodyParser(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}

	if masyarakat.Nama == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Nama required"})
	}

	if masyarakat.No_hp == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Nomor hp required"})
	}

	if masyarakat.Tempat_lahir == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Tempat lahir required"})
	}

	// if masyarakat.Birthday == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Tanggal lahir required"})
	// }

	if masyarakat.Alamat == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Alamat required"})
	}

	if err := models.ValidateMasyarakat(&masyarakat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg_validate": err.Error()})
	}

	if err := tx.Where("nik = ?", nik).Updates(&masyarakat).Error; err != nil {
		return c.Status(fiber.StatusNotModified).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{"msg": "Profile berhasil di update"})
}

func DeleteProfile(c *fiber.Ctx) error {
	nik := c.Params("nik")
	if nik == "" {
		return c.Status(404).JSON(fiber.Map{"msg": "NIK required"})
	}

	var user models.User
	if err := models.DB.Where("id = ?", nik).Delete(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{"msg": "User berhasil dihapus!"})
}
