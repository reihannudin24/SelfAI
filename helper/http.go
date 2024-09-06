package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func ReadFromRequestBody(r *http.Request, result interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		return err
	}
	return nil
}

func GetHeaderAuth(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return "", errors.New("Invalid authorization header format")
	}

	return tokenParts[1], nil
}
