package phone_model

type Phone struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Phone       string `json:"phone_model"`
	CreatedBy   string `json:"created_by"`
	CompanyName string `json:"company_name"`
}
