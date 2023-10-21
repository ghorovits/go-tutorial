package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upperCurrency := strings.ToUpper(currency)

// make an http request
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		// response is not a string or bytes, so we need to use this package to read it
		dataInBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		// passing the address of the reponse variable so it can modify it
		err = json.Unmarshal(dataInBytes, &response)

		if err != nil {
			return nil, err
		}


	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}