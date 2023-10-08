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

	ktpFile, err := c.FormFile("ktp_lama")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	akteFile, err := c.FormFile("akte")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	ijazahFile, err := c.FormFile("ijazah")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	var masyarakat models.Masyarakat
	if err := models.DB.Where("idm =? ", surat.Id_masyarakat).First(&masyarakat).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"msg": "Data tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"msg": err.Error()})
	}

	splitKtp := strings.Split(ktpFile.Filename, ".")
	splitIjazah := strings.Split(ijazahFile.Filename, ".")
	splitAkte := strings.Split(akteFile.Filename, ".")

	if splitKtp[1] != "pdf" || splitIjazah[1] != "pdf" || splitAkte[1] != "pdf" {
		return c.Status(400).JSON(fiber.Map{"msg": "File harus berkekstensi PDF"})
	}

	ktpName := masyarakat.NIK + ktp.ID + "-KTP.pdf"
	akteName := masyarakat.NIK + ktp.ID + "-AKTE.pdf"
	ijazahName := masyarakat.NIK + ktp.ID + "-IJAZAH.pdf"
	if err := c.SaveFile(ktpFile, fmt.Sprintf("./public/pengantarktp/ktp_lama/%s", ktpName)); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := c.SaveFile(akteFile, fmt.Sprintf("./public/pengantarktp/akte/%s", akteName)); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	if err := c.SaveFile(ijazahFile, fmt.Sprintf("./public/pengantarktp/ijazah/%s", ijazahName)); err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	ktp.Ktp_lama = ktpName
	ktp.Akte = akteName
	ktp.Ijazah = ijazahName
	if err := models.DB.Create(&ktp).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": err.Error()})
	}

	return c.JSON(fiber.Map{
		"msg": "Pengajuan berhasil",
		// "ucob": splitAkte[1],
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
	return nil
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
