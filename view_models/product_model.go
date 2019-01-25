package view_models

type ProductModel struct {
	ProductId   string `gorm:"primary_key" json:"product_id"`
	Name        string `gorm:"not null" json:"name"`
	ImageOpen   string `json:"image_open"`
	ImageClosed string `json:"image_closed"`
	Description string `json:"description"`

	Story       string `json:"story"`
	AllergyInfo string `json:"allergy_info"`

	DietaryCertifications string `json:"dietary_certifications"`

	SourcingValues []string `json:"sourcing_values"`
	Ingredients    []string `json:"ingredients"`
}
