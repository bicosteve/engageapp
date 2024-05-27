package controllers

import (
	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/utils"
	"net/http"
	"strconv"
)

func (b *Base) CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload entities.PostPayload
	err := utils.ReadJSON(w, r, payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = entities.ValidatePost(&payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	claims, err := entities.ValidateClaims(&entities.Claims{}, r)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(claims.ID)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = userId

}
