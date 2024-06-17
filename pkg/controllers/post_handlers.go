package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/models"
	"github.com/engageapp/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

func (b *Base) CreatePost(w http.ResponseWriter, r *http.Request) {
	var payload *entities.PostPayload

	err := utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	tknString, err := b.User.GetTokenString(r)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	token, err := b.User.ValidateClaim(tknString)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !token.Valid {
		utils.ErrorJSON(w, errors.New("invalid token"), http.StatusBadRequest)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	usr := claims["userId"].(string)

	userId, err := strconv.Atoi(usr)
	if err != nil {
		utils.ErrorJSON(w, errors.New(err.Error()), http.StatusBadRequest)
		return
	}

	err = models.CreatePost(payload, userId, b.DB)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"msg": "Post Created"})
}

func (b *Base) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetPosts(b.DB)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"posts": posts})
}
