package handlers

import (
	"encoding/json"
	"net/http"
	"rest_api_sklad_project/package/mocks"

	"github.com/gorilla/mux"
)


// AddGood Add good
//
//	@Summary		Add good
//	@Description	Add good
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			Id	path		string	true	"Goods Id for delete"
//	@Success		200	{string}	string
//	@Router			/goods/{id} [delete]
func DeleteGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, goods := range mocks.Goods {
		if goods.Id == id {
			mocks.Goods = append(mocks.Goods[:index], mocks.Goods[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
