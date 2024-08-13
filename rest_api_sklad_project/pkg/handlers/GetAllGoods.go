package handlers

import (
	"encoding/json"
	"net/http"
	"rest_api_sklad_project/package/mocks"
)

// GetAllGoods Get all goods
//
//	@Summary		Get all goods
//	@Description	Get all goods
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/goods [get]
func GetAllGoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Goods)
}
