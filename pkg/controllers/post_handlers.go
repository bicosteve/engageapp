package controllers

import (
	"net/http"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/models"
	"github.com/engageapp/pkg/utils"
)

func (b *Base) CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload *entities.PostPayload
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// err = entities.ValidatePost(payload)

	// if err != nil {
	// 	utils.ErrorJSON(w, err, http.StatusBadRequest)
	// 	return
	// }

	userId, err := models.ValidClaim(b.UserValidator, r)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	//err = b.PostModel.CreatePost(payload, userId, b.DB)

	err = models.CreatePost(payload, payload, userId, b.DB)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Post Created"})

}
