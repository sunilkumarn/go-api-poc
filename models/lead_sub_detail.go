package models

type LeadSubDetail struct {
	BaseModel
	Website      string `json:"website"`
	CallOrNot    string `gorm:"default:0" json:"call_or_not"`
	ReadOrNot    string `json:"read_or_not" gorm:"default:0"`
	Empct        string `json:"empct" gorm:"default:0"`
	LeadDetailID uint   `json:"-"`
}
