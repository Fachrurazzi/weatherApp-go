package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ToEncode(writer http.ResponseWriter, object interface{}) []byte {
	writer.Header().Add("Content-Type", "application/json")
	respose, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
	}
	return respose
}

func ToDecode(body []byte, object interface{}) error {
	err := json.Unmarshal(body, object)
	if err != nil {
		PanicIfError(err)
	}
	return nil
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
