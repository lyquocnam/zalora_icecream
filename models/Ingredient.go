package models

const IngredientTable = "ingredients"

type Ingredient struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`

	Products []Product `gorm:"many2many:products_ingredients" json:"products"`
}

func (Ingredient) TableName() string {
	return IngredientTable
}
