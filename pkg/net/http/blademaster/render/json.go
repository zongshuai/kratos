package render

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

// JSON common json struct.
type JSON struct {
	Code       int         `json:"error_code"`
	Message    string      `json:"error_message"`
	ServerTime int64       `json:"server_time"`
	Data       interface{} `json:"data,omitempty"`
	TraceId    string      `json:"trace_id"`
}

func writeJSON(w http.ResponseWriter, obj interface{}) (err error) {
	var jsonBytes []byte
	writeContentType(w, jsonContentType)
	if jsonBytes, err = json.Marshal(obj); err != nil {
		err = errors.WithStack(err)
		return
	}
	if _, err = w.Write(jsonBytes); err != nil {
		err = errors.WithStack(err)
	}
	return
}

// Render (JSON) writes data with json ContentType.
func (r JSON) Render(w http.ResponseWriter) error {
	// FIXME(zhoujiahui): the TTL field will be configurable in the future
	return writeJSON(w, r)
}

// WriteContentType write json ContentType.
func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// MapJSON common map json struct.
type MapJSON map[string]interface{}

// Render (MapJSON) writes data with json ContentType.
func (m MapJSON) Render(w http.ResponseWriter) error {
	return writeJSON(w, m)
}

// WriteContentType write json ContentType.
func (m MapJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
