package utilify

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strings"
)

const (
	contentType     = "content-type"
	contentTypeJson = "application/json"
)

var (
	ErrEOF           = errors.New("body must not be empty")
	ErrInvalidJson   = errors.New("body contains badly-formed JSON")
	ErrDuplicateJson = errors.New("body contains only one JSON object")
	ErrNotJson       = errors.New("body content-type header is not application/json")
)

func ReadJson(r *http.Request, dst interface{}) error {
	contentType := r.Header.Get(contentType)
	if contentType != "" && strings.ToLower(contentType) != contentTypeJson {
		return ErrNotJson
	}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	err := d.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return ErrInvalidJson
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return ErrEOF
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	if err = d.Decode(&struct{}{}); err != io.EOF {
		return ErrDuplicateJson
	}
	v := validator.New()
	return v.Struct(dst)
}

func WriteJson(w http.ResponseWriter, status int, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set(contentType, contentTypeJson)
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}
