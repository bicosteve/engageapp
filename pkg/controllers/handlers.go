package controllers

// Contains code for RabbitMQ producer

import (
	"net/http"

	"github.com/engageapp/pkg/entities"
	"github.com/engageapp/pkg/utils"
)

type JSONResponse utils.JSONResponse

// Broker() -> Broker handle
func (b *Base) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JSONResponse{
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
	err = entities.Validate(userRequestBody)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"msg": "User Created!"})

}
