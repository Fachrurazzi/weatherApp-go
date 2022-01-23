package helper

import (
	"io/ioutil"
	"net/http"
	"time"
)

func MakeApiRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
