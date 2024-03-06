package httpkit

import (
	"encoding/json"
	"github.com/trangnkp/my_books/src/container"
	"log"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Verdict    string      `json:"verdict"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (r *Response) IsEmpty() bool {
	return r.StatusCode == 0 || len(r.Verdict) == 0
}

func SendJSON(
	w http.ResponseWriter,
	statusCode int, verdict, message string, data interface{}) error {
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(statusCode)

	obj := container.Map{
		"verdict": verdict,
		"message": message,
		"data":    data,
		"time":    time.Now().Format(DateTimeLayout),
	}
	body, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
