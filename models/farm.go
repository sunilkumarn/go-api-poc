package models

type Farm struct {
	BaseModel
	AllOrganic bool   `json:"all_organic" gorm:"default:0"`
	Website    string `json:"website"`
	OwnerName  string `json:"owner_name"`
	AllInHouse bool   `json:"all_in_house" gorm:"default:0"`
}
