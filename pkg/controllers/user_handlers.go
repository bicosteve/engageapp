package controllers

// Contains code for RabbitMQ producer

import (
	"net/http"
	"time"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/utils"
)

// Broker() -> Broker handle
func (b *Base) Broker(w http.ResponseWriter, r *http.Request) {
	payload := entities.JSONResponse{
		Error:   false,
		Message: "Hit Broker",
	}

	err := utils.WriteJSON(w, http.StatusOK, payload)

	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

}

// PostUser() -> checks user request body and validates it
func (b *Base) PostUser(w http.ResponseWriter, r *http.Request) {
	userRequestBody := new(entities.UserPayload)

	err := utils.ReadJSON(w, r, userRequestBody)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Validate the payload
	err = entities.ValidateUser(userRequestBody)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = b.UserModel.RegisterUser(userRequestBody, b.DB)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.PublishToQueue(b.Chan, "Test", userRequestBody)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "User Created!"})

}

func (b *Base) Login(w http.ResponseWriter, r *http.Request) {
	payload := new(entities.UserPayload)

	err := utils.ReadJSON(w, r, payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = entities.ValidateLogins(payload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return

	}

	token, err := b.UserModel.LoginUser(payload, b.DB)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.PublishToQueue(b.Chan, "Test", token)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		SameSite: 4,
		Secure:   true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	http.SetCookie(w, cookie)
	r.AddCookie(cookie)

	// utils.WriteJSON(w, http.StatusAccepted, map[string]string{"token": token})
}