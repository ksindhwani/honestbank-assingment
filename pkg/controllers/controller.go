package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/honestbank/tech-assignment-backend-engineer/pkg/httputils"
	"github.com/honestbank/tech-assignment-backend-engineer/pkg/service"
)

func ProcessData(resp http.ResponseWriter, req *http.Request, validator *validator.Validate) {
	switch req.Method {
	case http.MethodPost:
		checkCreditStatus(resp, req, validator)
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}

}

func checkCreditStatus(resp http.ResponseWriter, req *http.Request, validator *validator.Validate) {
	var t service.RecordData
	body, err := httputils.GetRequestBody(resp, req)
	if err != nil {
		httputils.WriteErrorResponse(resp, httputils.NewBadRequestError(err, "Unable to fetch request body"))
		return
	}
	if err := json.Unmarshal(body, &t); err != nil {
		httputils.WriteErrorResponse(resp, httputils.NewBadRequestError(err, "Unable to parse request body"))
		return
	}

	if err = validator.Struct(t); err != nil {
		httputils.WriteErrorResponse(resp, httputils.NewBadRequestError(err, "Required fields missing in request body"))
		return
	}

	response, err := service.ProcessCreditCardApplication(t)
	if err != nil {
		httputils.WriteErrorResponse(resp, httputils.NewInternalServerError(err, "Unable to process application"))
		return
	}
	httputils.WriteResponse(resp, http.StatusOK, response)
}
