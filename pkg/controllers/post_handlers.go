package controllers

import (
	"net/http"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/utils"
)

func (b *Base) CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload *entities.PostPayload
	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = entities.ValidatePost(payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId, err := entities.ValidateClaims(r)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// userId, err := strconv.Atoi(claims.ID)
	// if err != nil {
	// 	utils.ErrorJSON(w, err, http.StatusBadRequest)
	// 	return
	// }

	// fmt.Println(userId)

	err = b.PostModel.CreatePost(payload, userId, b.DB)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Post Created"})

}
