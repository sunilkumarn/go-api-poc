package models

type FarmAddress struct {
	BaseModel
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	Pobox     string `json:"pobox"`
	Mobile    string `json:"mobile"`
	CallOrNot bool   `json:"call_or_not" gorm:"default:0"`
	Farm      Farm   `json:"-"`
}
