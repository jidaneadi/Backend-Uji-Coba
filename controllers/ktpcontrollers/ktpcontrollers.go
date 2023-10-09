package ktpcontrollers

import (
	"Backend_TA/models"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	var surat models.Surat
	if err := c.BodyParser(&surat); err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
	}

	if surat.Id_masyarakat == "" {
		return c.Status(404).JSON(fiber.Map{"msg": "Data masyarakat kosong"})
	}

	if err := models.DB.Create(&surat).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	var ktp models.Pengantar_KTP
	ktp.Id_surat = surat.ID
	if err := c.BodyParser(&ktp); err != nil {
		return c.Status(404).JSON(fiber.Map{"msg": err.Error()})
	}

	dokumenSyarat, err := c.FormFile("dokumen")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	//Mengambil data NIK masyarakatcls
	var masyarakat models.Masyarakat
	if err := models.DB.Where("idm =? ", surat.Id_masyarakat).First(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}

	cekFormat := strings.Split(dokumenSyarat.Filename, ".")

	if cekFormat[1] != "pdf" {
		return c.Status(400).JSON(fiber.Map{"msg": "File harus berkekstensi PDF"})
	}
	namaDokumen := masyarakat.NIK + ktp.Id_surat + "-KTP.pdf"
	if err := c.SaveFile(dokumenSyarat, fmt.Sprintf("./public/pengantarktp/%s", namaDokumen)); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})

	}

	ktp.Dokumen_syarat = namaDokumen
	if err := models.DB.Create(&ktp).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{
		"msg": "Pengajuan berhasil",
	})
}

func Show(c *fiber.Ctx) error {
	var ktp []models.Pengantar_KTP

	//Join 3 tabel
	if err := models.DB.Preload("Surat.Masyarakat").Find(&ktp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "Data null"})
		}
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}

	data := make([]fiber.Map, len(ktp))
	for i, dataKtp := range ktp {
		data[i] = fiber.Map{
			"id_surat":   dataKtp.Id_surat,
			"nik":        dataKtp.Surat.Masyarakat.NIK,
			"nama":       dataKtp.Surat.Masyarakat.Nama,
			"syarat":     dataKtp.Dokumen_syarat,
			"jns_surat":  dataKtp.Surat.Jns_surat,
			"status":     dataKtp.Surat.Status,
			"tgl":        dataKtp.Surat.UpdatedAt.String()[0:10],
			"keterangan": dataKtp.Surat.Keterangan,
		}
	}
	return c.JSON(data)
}

func ShowId(c *fiber.Ctx) error {
	id := c.Params("id")
	var pengantar_ktp models.Pengantar_KTP

	if err := models.DB.Preload("Surat.Masyarakat").
		Joins("JOIN surat ON surat.id = pengantar_ktp.id_surat").
		Joins("JOIN masyarakat ON masyarakat.idm = surat.id_masyarakat").
		Where("masyarakat.nik =?", id).
		First(&pengantar_ktp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "Data not found"})
		}
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}
	return c.JSON(fiber.Map{
		"id_surat":   pengantar_ktp.Id_surat,
		"NIK":        pengantar_ktp.Surat.Masyarakat.NIK,
		"Nama":       pengantar_ktp.Surat.Masyarakat.Nama,
		"syarat":     pengantar_ktp.Dokumen_syarat,
		"jns_surat":  pengantar_ktp.Surat.Jns_surat,
		"status":     pengantar_ktp.Surat.Status,
		"tgl":        pengantar_ktp.Surat.UpdatedAt.String()[0:10],
		"keterangan": pengantar_ktp.Surat.Keterangan,
	})
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
