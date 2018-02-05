package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadJSON(body io.Reader, v interface{}) error {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(byt, v)
}

func WriteJSON(body io.Writer, v interface{}) error {
	byt, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := body.Write(byt); err != nil {
		return err
	}
	return nil
}

func JSONBuffer(v interface{}) (*bytes.Buffer, error) {
	byt, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(byt), nil
}
