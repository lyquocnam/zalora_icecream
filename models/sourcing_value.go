package models

const SourcingValueTable = "sourcing_values"

type SourcingValue struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`

	Products []Product `gorm:"many2many:products_sourcing_values" json:"products"`
}

func (SourcingValue) TableName() string {
	return SourcingValueTable
}
