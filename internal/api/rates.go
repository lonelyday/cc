package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lonelyday/cc/pkg/model"

	"github.com/lonelyday/cc/internal/httpclient"

	"github.com/gin-gonic/gin"
)

func Rates(c *gin.Context) {

	currencies := strings.Split(c.Request.URL.Query().Get("currencies"), ",")
	if len(currencies) < 2 {
		log.Println("Not enough currencies to convert between provided")
		c.Status(http.StatusBadRequest)
		return
	}
	r, err := getExchangeRates(currencies)
	if err != nil {
		log.Println("Error fetching exchange rates:", err)
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, r)
}

func findRates(currencies []string, d model.OpenCurrencyConverter) ([]model.Rates, error) {
	var r []model.Rates

	for i := 0; i < len(currencies)-1; i++ {
		for j := i + 1; j < len(currencies); j++ {
			from, f_ok := d.Rates[currencies[i]]
			to, t_ok := d.Rates[currencies[j]]
			r = append(r, model.Rates{From: currencies[i], To: currencies[j]})
			r = append(r, model.Rates{From: currencies[j], To: currencies[i]})
			if f_ok && t_ok {
				r[len(r)-2].Rate = from / to
				r[len(r)-1].Rate = to / from
			} else {
				return nil, fmt.Errorf("One of the currencies (%s, %s) not found", currencies[i], currencies[j])
			}
		}
	}
	return r, nil

}

func getExchangeRates(currencies []string) ([]model.Rates, error) {
	client := httpclient.Client()
	req, err := httpclient.OERReq()
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var d model.OpenCurrencyConverter
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return nil, err
	}
	return findRates(currencies, d)

}
