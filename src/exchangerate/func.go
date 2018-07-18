package exchangerate

import (
	"html/template"
	"net/http"
)

func AddUnitView() (*template.Template, interface{}, error) {
	tmpl, data, err := addUnitView()
	if err != nil {
		return nil, nil, err
	}

	return tmpl, data, nil
}

func AddUnit(r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		return false
	}

	unit := r.Form.Get("unit")
	description := r.Form.Get("description")

	result, err := insertUnit(unit, description)
	if err != nil {
		panic(err)
	}

	return result
}

func AddExchangeView() (*template.Template, map[string][]Currency, error) {
	unit, err := getAllUnit()
	if err != nil {
		return nil, nil, err
	}

	tmpl, data, err := addExchangeView(unit)
	if err != nil {
		return nil, nil, err
	}

	return tmpl, data, nil
}

func RemoveExchangeView() (*template.Template, map[string][]CurrencyExchange, error) {
	exchange, err := getAllExchange()
	if err != nil {
		return nil, nil, err
	}

	tmpl, data, err := removeExchangeView(exchange)
	if err != nil {
		return nil, nil, err
	}

	return tmpl, data, nil
}

func AddExchangeRateView() (*template.Template, map[string][]CurrencyExchange, error) {
	exchange, err := getAllExchange()
	if err != nil {
		return nil, nil, err
	}

	tmpl, data, err := addExchangeRateView(exchange)
	if err != nil {
		return nil, nil, err
	}

	return tmpl, data, nil
}

func ViewExchangeRateView(r *http.Request) (*template.Template, map[string][]CurrencyExchangeRate, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, nil, err
	}

	date := r.Form.Get("date")
	var exchangeRate []CurrencyExchangeRate
	exchangeRate = nil

	if date != "" {
		var err error
		exchangeRate, err = getExchangeHistory(date)
		if err != nil {
			return nil, nil, err
		}
	}

	tmpl, data, err := viewExchangeRateView(exchangeRate)
	if err != nil {
		return nil, nil, err
	}

	return tmpl, data, nil
}

func GetUnit(r *http.Request) ([]Currency, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	id_unit_exclude := r.Form.Get("exclude")
	unit, err := getUnitExclude(id_unit_exclude)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func AddExchange(r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		return false
	}

	idUnitFrom := r.Form.Get("unit_from")
	idUnitTo := r.Form.Get("unit_to")

	result, err := insertExchange(idUnitFrom, idUnitTo)
	if err != nil {
		panic(err)
	}

	return result
}

func RemoveExchange(r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		return false
	}

	idExchange := r.Form.Get("id_exchange")

	result, err := removeExchange(idExchange)
	if err != nil {
		panic(err)
	}

	return result
}

func AddExchangeRate(r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		return false
	}

	date := r.Form.Get("date")
	idExchange := r.Form.Get("exchange")
	rate := r.Form.Get("rate")

	result, err := insertExchangeRate(date, idExchange, rate)
	if err != nil {
		panic(err)
	}

	return result
}
