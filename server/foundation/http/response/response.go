package response

import (
	"encoding/json"
	"net/http"
)


func Write(writer http.ResponseWriter, data any, statusCode int) (int, error) {
	var err error
	var dataToWrite []byte

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if data == nil {
		return writer.Write(nil)
	}

	dataToWrite, err = json.Marshal(data)
	if err != nil {
		return 0, err
	}

	return writer.Write(dataToWrite)
}

