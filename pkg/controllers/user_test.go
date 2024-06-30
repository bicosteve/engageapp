package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var base Base

func TestCreateUser(t *testing.T) {
	payload := []byte(`{"email":"bico.steve@gmail.com","password":"1234","confirm_password":"1234"}`)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder() // rr-> response recorder
	handler := http.HandlerFunc(base.PostUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

}

func TestInvalidCredentials(t *testing.T) {
	// Simulating invalid json
	payload := []byte(`{"email":"bico.steve@gmail.com","password":"1234"}`)
	request, err := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(base.PostUser)

	handler.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
