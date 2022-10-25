package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

const defaultMaxBodyBytes int64 = 1048576 // 1MB

type UnsupportedMediaType struct {
	Err error `json:"error"`
}

func (err *UnsupportedMediaType) Error() string {
	return err.Err.Error()
}

func Decode(r *http.Request, dest interface{}, maxBodyBytes ...int64) error {
	maxBody := defaultMaxBodyBytes
	if maxBodyBytes != nil && maxBodyBytes[0] > 0 {
		maxBody = maxBodyBytes[0]
	}

	r.Body = http.MaxBytesReader(nil, r.Body, maxBody)

	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		return decodeJson(r, dest)
	default:
		return &UnsupportedMediaType{Err: errors.New("invalid content type")}
	}
}

// decodeJson decodes from json body
func decodeJson(r *http.Request, dest interface{}) error {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(dest); err != nil {
		return err
	}

	return nil
}
