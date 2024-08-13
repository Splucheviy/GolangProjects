package handlers

import (
	"encoding/json"
	"net/http"
	"rest_api_sklad_project/package/mocks"

	"github.com/gorilla/mux"
)

// GetGood Get goods by id
//
//	@Summary		Get goods by id
//	@Description	Get goods by id
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			Id	path		string	true	"Goods id"
//	@Success		200	{string}	string
//	@Router			/goods/{Id} [get]
func GetGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, goods := range mocks.Goods {
		if goods.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(goods)
			break
		}
	}
}
