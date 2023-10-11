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
	var dataMasyarakat []models.Masyarakat

	//Join 3 tabel
	if err := models.DB.Preload("Surat.Pengantar_KTP").
		Joins("JOIN surat ON surat.id_masyarakat = masyarakat.idm").
		Joins("JOIN pengantar_ktp ON pengantar_ktp.id_surat = surat.id").
		Find(&dataMasyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "Data null"})
		}
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}

	data := make([]fiber.Map, len(dataMasyarakat))
	for i, dataKtp := range dataMasyarakat {
		data[i] = fiber.Map{
			"id_surat":   dataKtp.Surat.ID,
			"nik":        dataKtp.NIK,
			"nama":       dataKtp.Nama,
			"syarat":     dataKtp.Surat.Pengantar_KTP.Dokumen_syarat,
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
	var dataMasyarakat []models.Masyarakat

	if err := models.DB.Preload("Surat.Pengantar_KTP").
		Where("masyarakat.nik =?", id).
		Find(&dataMasyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "Data not found"})
		}
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}
	data := make([]fiber.Map, len(dataMasyarakat))
	for i, dataKtp := range dataMasyarakat {
		data[i] = fiber.Map{
			"id_surat":   dataKtp.Surat.ID,
			"nik":        dataKtp.NIK,
			"nama":       dataKtp.Nama,
			"syarat":     dataKtp.Surat.Pengantar_KTP.Dokumen_syarat,
			"jns_surat":  dataKtp.Surat.Jns_surat,
			"status":     dataKtp.Surat.Status,
			"tgl":        dataKtp.Surat.UpdatedAt.String()[0:10],
			"keterangan": dataKtp.Surat.Keterangan,
		}
	}
	return c.JSON(data)
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
