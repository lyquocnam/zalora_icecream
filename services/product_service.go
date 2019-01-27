package services

import (
	"github.com/lyquocnam/zalora_icecream/lib"
	"github.com/lyquocnam/zalora_icecream/models"
	"github.com/lyquocnam/zalora_icecream/view_models"
	"strconv"
)

func FindOneProductById(id string) *models.Product {
	result := &models.Product{}

	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		First(&result, id)

	return result
}

func FindOneProductModelById(id string) *view_models.ProductModel {
	var product models.Product
	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		First(&product, id)

	if len(product.ID) < 1 {
		return nil
	}

	result := product.ToModel()

	return &result
}

func ProductExistByName(name string) bool {
	count := 0
	lib.DB.Model(&models.Product{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func ProductSourcingExist(productId string, sourcingId int) bool {
	count := 0
	lib.DB.
		Model(&models.ProductSourcingValue{}).
		Where("product_id = ? and sourcing_value_id = ?", productId, sourcingId).
		Count(&count)
	return count > 0
}

func ProductIngredientExist(productId string, ingredientId int) bool {
	count := 0
	lib.DB.
		Model(&models.ProductIngredient{}).
		Where("product_id = ? and ingredient_id = ?", productId, ingredientId).
		Count(&count)
	return count > 0
}

func ConvertListStringToListInt(values []string) ([]int, error) {
	result := make([]int, 0)

	if len(values) < 1 {
		return result, nil
	}

	for _, value := range values {
		v, err := strconv.Atoi(value)
		if err != nil {
			return result, err
		}
		result = append(result, v)
	}

	return result, nil
}

func ProductContainSourcingValues(sourcingValueIds []int) bool {
	count := 0
	lib.DB.Model(&models.SourcingValue{}).Where("id in (?)", sourcingValueIds).Count(&count)
	return count > 0
}

func ProductContainIngredients(ingredientIds []int) bool {
	count := 0
	lib.DB.Model(&models.Ingredient{}).Where("id in (?)", ingredientIds).Count(&count)
	return count > 0
}

func FindProduct() (result []models.Product) {
	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		First(&result)

	return result
}

func FindProductModel() (result []view_models.ProductModel) {
	var products []models.Product
	lib.DB.
		Preload("Ingredients").
		Preload("SourcingValues").
		First(&products)

	for _, p := range products {
		result = append(result, p.ToModel())
	}

	return result
}
