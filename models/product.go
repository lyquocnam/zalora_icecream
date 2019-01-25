package models

const ProductTable = "products"

type Product struct {
	ProductId   string `gorm:"primary_key" json:"product_id"`
	Name        string `gorm:"not null" json:"name"`
	ImageOpen   string `json:"image_open"`
	ImageClosed string `json:"image_closed"`
	Description string `json:"description"`

	Story       string `gorm:"not null" json:"story"`
	AllergyInfo string `gorm:"not null" json:"allergy_info"`

	DietaryCertifications string `json:"dietary_certifications"`

	SourcingValues []SourcingValue `gorm:"many2many:products_sourcing_values" json:"sourcing_values"`
	Ingredients    []Ingredient    `gorm:"many2many:products_ingredients" json:"ingredients"`
}

func (Product) TableName() string {
	return ProductTable
}
