package mocks

import "rest_api_sklad_project/package/models"

var Goods = []models.Goods{
	{
		Id: "c5198c5a-44ab-45ab-88a7-b540825b946c", 
		Name: "Good 1", 
		Provider: 1, 
		Price: 14, 
		Quantity: 25},
	{
		Id: "a35f00e6-1baa-4023-a4a3-501cc8a531e3", 
		Name: "Good 2", 
		Provider: 2,
		Price: 24, 
		Quantity: 100},
}
