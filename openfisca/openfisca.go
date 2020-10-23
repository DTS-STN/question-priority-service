package openfisca

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Trace(data []byte) (body []byte, err error) {

	req, err := http.NewRequest("POST", "http://localhost:5000/trace", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error in req: ", err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in request: ", err)
	}

	// Defer the closing of the body
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Printf("Error reading body: %v", err)
		//http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	return body, nil
}