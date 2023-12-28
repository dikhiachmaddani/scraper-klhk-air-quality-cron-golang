package models

type QualityAir struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Pm10      int    `json:"pm10"`
	Pm25      int    `json:"pm25"`
	So2       int    `json:"so2"`
	Co        int    `json:"co"`
	O3        int    `json:"o3"`
	No2       int    `json:"no2"`
	Hc        int    `json:"hc"`
	Station   string `json:"id_stasiun"`
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
	Waktu     string `json:"waktu"`
	Ispu      int    `json:"ispu"`
	Category  string `json:"category"`
}

type Rows struct {
	Rows []GetDataQuality `json:"rows"`
}

type GetDataQuality struct {
	Pm10      string `json:"pm10"`
	Pm25      string `json:"pm25"`
	So2       string `json:"so2"`
	Co        string `json:"co"`
	O3        string `json:"o3"`
	No2       string `json:"no2"`
	Hc        string `json:"hc"`
	Station   string `json:"id_stasiun"`
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
	Waktu     string `json:"waktu"`
	Val       string `json:"val"`
	Cat       string `json:"cat"`
}
