package handlers

import (
	"app/interfaces/errs"
	"encoding/json"
	"net/http"
)

// decodeR decodes request's body to given interface
func decodeR(r *http.Request, to interface{}) error {
	return errs.Wrap(json.NewDecoder(r.Body).Decode(to))
}

type response struct {
	Result interface{} `json:"result"`
}
