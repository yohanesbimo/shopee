package exchangerate

import (
	"html/template"
	"os"
)

func viewPath(view string) string {
	pwd, _ := os.Getwd()
	return pwd + "/src/exchangerate/views/" + view + ".html"
}

func addUnitView() (*template.Template, interface{}, error) {
	tmpl, err := template.ParseFiles(viewPath("addUnit"))
	if err != nil {
		return nil, nil, err
	}

	return tmpl, nil, nil
}

func addExchangeView(unit []Currency) (*template.Template, map[string][]Currency, error) {
	tmpl, err := template.ParseFiles(viewPath("addExchange"))
	if err != nil {
		return nil, nil, err
	}

	data := map[string][]Currency{
		"Unit": unit,
	}

	return tmpl, data, nil
}

func removeExchangeView(exchange []CurrencyExchange) (*template.Template, map[string][]CurrencyExchange, error) {
	tmpl, err := template.ParseFiles(viewPath("removeExchange"))
	if err != nil {
		return nil, nil, err
	}

	data := map[string][]CurrencyExchange{
		"Exchange": exchange,
	}

	return tmpl, data, nil
}

func addExchangeRateView(exchange []CurrencyExchange) (*template.Template, map[string][]CurrencyExchange, error) {
	tmpl, err := template.ParseFiles(viewPath("addExchangeRate"))
	if err != nil {
		return nil, nil, err
	}

	data := map[string][]CurrencyExchange{
		"Exchange": exchange,
	}

	return tmpl, data, nil
}

func viewExchangeRateView(currencyExchangeRate []CurrencyExchangeRate) (*template.Template, map[string][]CurrencyExchangeRate, error) {
	tmpl, err := template.ParseFiles(viewPath("viewExchangeRate"))
	if err != nil {
		return nil, nil, err
	}

	data := map[string][]CurrencyExchangeRate{
		"ExchangeRate": currencyExchangeRate,
	}

	return tmpl, data, nil
}
