package model

import (
	"time"
)

type (
	Carfix struct {
		VehicleCode   string `json:"VehicleCode"`
		KodeBrand     string `json:"KodeBrand"`
		BrandMobil    string `json:"BrandMobil"`
		KodeModel     string `json:"KodeModel"`
		ModelMobil    string `json:"ModelMobil"`
		TipeMobil     string `json:"TipeMobil"`
		VehicleTypeID string `json:"VehicleTypeId"`
		TahunMobil    string `json:"TahunMobil"`
		KategoriMobil string `json:"KategoriMobil"`
		GradeMobil    string `json:"GradeMobil"`
		MesinMobile   string `json:"MesinMobile"`
		VehicleID     string `json:"VehicleId"`
		TglUpload     string `json:"TglUpload"`
		Area          string `json:"Area"`
	}

	User struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Bayauct struct {
		IDBayauct   string    `json:"idBayauct"`
		VehicleCode string    `json:"VehicleCode"`
		Merk        string    `json:"Merk"`
		Kategori    string    `json:"Kategori"`
		Model       string    `json:"Model"`
		Tahun       string    `json:"Tahun"`
		Warna       string    `json:"Warna"`
		NoPolisi    string    `json:"NoPolisi"`
		HargaLimit  uint      `json:"HargaLimit"`
		HargaSukses uint      `json:"HargaSukses"`
		Result      string    `json:"Result"`
		TglUpload   time.Time `json:"TglUpload" time:"2006-01-02 15:04:05"`
		Area        string    `json:"Area"`
	}
)

func (Carfix) TableName() string {
	return "tb_carfix"
}

func (User) TableName() string {
	return "tb_users"
}

func (Bayauct) TableName() string {
	return "bay_sync"
}
