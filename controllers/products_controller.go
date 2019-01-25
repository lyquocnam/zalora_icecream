package controllers

import (
	"github.com/labstack/echo"
	"github.com/lyquocnam/zalora_icecream/lib"
	"github.com/lyquocnam/zalora_icecream/models"
	"github.com/lyquocnam/zalora_icecream/view_models"
	"net/http"
)

func ProductGetListHandler(c echo.Context) error {
	var products = make([]models.Product, 0)
	lib.DB.Preload("Ingredients").Preload("SourcingValues").Find(&products)
	result := make([]view_models.ProductModel, 0)

	for _, item := range products {
		ingredients := make([]string, 0)
		for _, ingredient := range item.Ingredients {
			ingredients = append(ingredients, ingredient.Name)
		}

		sourcingValues := make([]string, 0)
		for _, sourcingValue := range item.SourcingValues {
			sourcingValues = append(sourcingValues, sourcingValue.Name)
		}

		result = append(result, view_models.ProductModel{
			ProductId:             item.ProductId,
			Name:                  item.Name,
			Story:                 item.Story,
			AllergyInfo:           item.AllergyInfo,
			Description:           item.Description,
			ImageClosed:           item.ImageClosed,
			ImageOpen:             item.ImageOpen,
			DietaryCertifications: item.DietaryCertifications,
			Ingredients:           ingredients,
			SourcingValues:        sourcingValues,
		})
	}

	return c.JSON(http.StatusOK, &result)
}
