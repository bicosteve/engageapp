package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/models"
	"github.com/engageapp/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func (b *Base) CreateComment(w http.ResponseWriter, r *http.Request) {
	var payload *entities.CommentPayload

	postId, err := strconv.Atoi(chi.URLParam(r, "postid"))
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.ReadJSON(w, r, &payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	tknString, err := b.User.GetTokenString(r)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	tkn, err := b.User.ValidateClaim(tknString)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		utils.ErrorJSON(w, errors.New("invalid token"), http.StatusBadRequest)
		return
	}

	claims := tkn.Claims.(jwt.MapClaims)
	usr := claims["userId"].(string)

	userId, err := strconv.Atoi(usr)
	if err != nil {
		utils.ErrorJSON(w, errors.New(err.Error()), http.StatusBadRequest)
		return
	}

	err = models.AddComment(payload, userId, postId, b.DB)
	if err != nil {
		utils.ErrorJSON(w, errors.New(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{"msg": "Created"})
}

func (b *Base) GetComments(w http.ResponseWriter, r *http.Request) {
	postId, err := strconv.Atoi(chi.URLParam(r, "postid"))
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	comments, err := models.GetComment(postId, b.DB)
	if err != nil {
		utils.ErrorJSON(w, errors.New(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{"data": comments})
}
