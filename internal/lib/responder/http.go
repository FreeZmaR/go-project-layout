package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTP struct {
	w http.ResponseWriter
}

func NewHTTP(w http.ResponseWriter) *HTTP {
	return &HTTP{w: w}
}

func (r *HTTP) SendJSON(code int, body any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return r.Send(code, b)
}

func (r *HTTP) Send(code int, body []byte) error {
	r.w.WriteHeader(code)
	if _, err := r.w.Write(body); err != nil {
		return fmt.Errorf("error write response: %w", err)
	}

	return nil
}
