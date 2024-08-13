package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rest_api_sklad_project/package/mocks"
	"rest_api_sklad_project/package/models"

	"github.com/gorilla/mux"
)

// AddGood Add good
//
//	@Summary		Add good
//	@Description	Add good
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			Id			path		string	true	"Goods id to update"
//	@Param			Name		body		string	true	"Goods name to update"
//	@Param			Provider	body		integer	true	"Goods provider to update"
//	@Param			Price		body		integer	true	"Goods price to update"
//	@Param			Quantity	body		integer	true	"Goods quantity to update"
//	@Success		200			{string}	string
//	@Router			/goods/{id} [put]
func UpdateGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedGoods models.Goods
	json.Unmarshal(body, &updatedGoods)

	for index, goods := range mocks.Goods {
		if goods.Id == id {
			goods.Name = updatedGoods.Name
			goods.Provider = updatedGoods.Provider
			goods.Price = updatedGoods.Price
			goods.Quantity = updatedGoods.Quantity

			mocks.Goods[index] = goods

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
