package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Surat struct {
	ID            string         `gorm:"primariKey;autoIncrement" json:"id"`
	Id_masyarakat string         `json:"id_masyarakat" validate:"numeric,min=15"`
	Jns_surat     string         `gorm:"default:ktp" json:"jns_surat"`
	Status        string         `gorm:"default:diproses" json:"status"`
	Keterangan    string         `json:"keterangan"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Pengantar_KTP *Pengantar_KTP `gorm:"foreignKey:Id_surat;references:ID" json:"pengantar_ktp"`
}

func ValidateSurat(surat *Surat) error {
	validate := validator.New()
	return validate.Struct(surat)
}

func (Surat) TableName() string {
	return "surat"
}
