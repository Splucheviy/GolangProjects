package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"rest_api_sklad_project/package/mocks"
	"rest_api_sklad_project/package/models"

	"github.com/google/uuid"
)

// AddGood Add good
//
//	@Summary		Add good
//	@Description	Add good
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.Goods	true	"Goods data"
//	@Success		200		{string}	string
//	@Router			/goods [post]
func AddGood(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var product models.Goods
	json.Unmarshal(body, &product)

	product.Id = (uuid.New()).String()
	mocks.Goods = append(mocks.Goods, product)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
