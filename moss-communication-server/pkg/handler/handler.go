package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/MattiasHenders/moss-communication-server/pkg/errors"
	"github.com/go-chi/chi"
)

func Handler(h func(w http.ResponseWriter, r *http.Request) *errors.HTTPError) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		httpError := h(w, r)
		if httpError == nil {
			return
		}

		body, err := httpError.ResponseBody()
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(httpError.Status)
		w.Write(body)
	})
}

func GetURLParam(r *http.Request, field string) *string {
	param := chi.URLParam(r, field)
	if param == "" {
		return nil
	}

	return &param
}

func GetFormParam(r *http.Request, field string) *string {
	param := r.Form.Get(field)
	if strings.TrimSpace(param) == "" {
		return nil
	}

	return &param
}

func GetFormParamInt(r *http.Request, field string) *int {
	param := r.Form.Get(field)
	if param == "" {
		return nil
	}

	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return nil
	}

	return &paramInt
}

func GetFormParamBool(r *http.Request, field string) *bool {
	param := r.Form.Get(field)
	if param == "" {
		return nil
	}

	paramBool, err := strconv.ParseBool(param)
	if err != nil {
		return nil
	}

	return &paramBool
}

func GetQueryParam(r *http.Request, field string) *string {
	param := r.URL.Query().Get(field)
	if param == "" {
		return nil
	}

	return &param
}
