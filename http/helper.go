package httper

import (
	"encoding/json"
	"errors"
	"fmt"
	"medods/logger"
	"medods/services"
	"medods/utils"
	"net/http"
	"reflect"

	log "github.com/sirupsen/logrus"
)

type Helper struct {
	R       *http.Request
	W       http.ResponseWriter
	Service any
}

// CreateHandler - приводит функцию к http.HanlderFunc, позволяя в хендлерах использовать Helper.
// Так же служит единной точкой обратки респонсов.
func CreateHandler(fn func(Helper) (any, error), service any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := fn(Helper{R: r, W: w, Service: service})
		if err != nil {
			handleError(r, w, err)
		} else {
			handleResult(w, result)
		}
	}
}

func (h Helper) GetQueryParamAsString(name string) (string, error) {
	param := h.R.URL.Query().Get(name)

	if param == "" {
		return "", NewHttpErrorBadRequest(fmt.Sprintf("query param %v is empty", name))
	}

	return param, nil
}

// getErrorCode - возвращает http статус в зависимости от err.	
func getErrorCode(err error) int {
	if errors.Is(err, services.ErrTokenNotExist) {
		return http.StatusBadRequest
	}

	if httpError, ok := err.(HttpCodeError); ok {
		return httpError.Code
	}

	return http.StatusInternalServerError
}

func handleError(r *http.Request, w http.ResponseWriter, err error) {
	logger.Log().WithFields(log.Fields{"method": r.Method, "url": r.URL.Path}).Error(err)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(getErrorCode(err))

	b, err := json.Marshal(map[string]string{"error": err.Error()})
	if err != nil {
		log.Errorf("error on convert error to JSON - %s", err)
		return
	}

	if _, err := w.Write([]byte(b)); err != nil {
		log.Errorf("error on write error to response - %s", err)
	}
}

func handleResult(w http.ResponseWriter, result any) {
	if result == nil || reflect.ValueOf(result).IsNil() {
		w.WriteHeader(http.StatusOK)
		return
	}

	utils.WriteJson(w, result, http.StatusOK)
}
