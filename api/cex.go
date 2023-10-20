package api

import (
	"fmt"
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

	if res.StatusCode == http.StatusOK {
		
	}
}