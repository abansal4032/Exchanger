package common

import (
	"encoding/json"
	"net/http"
)

func ObjectToJson(object interface{}) ([]byte, error) {
	returnObject, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	return returnObject, nil
}

func WriteResponse(res http.ResponseWriter, response interface{}) {
	responseJson, err := ObjectToJson(response)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(responseJson)
}
