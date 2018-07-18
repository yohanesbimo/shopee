package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	er "exchangerate"
)

func renderHTML(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addUnitView(w http.ResponseWriter, r *http.Request) {
	tmpl, data, err := er.AddUnitView()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderHTML(w, tmpl, &data)
}

func actionAddUnit(w http.ResponseWriter, r *http.Request) {
	data := er.AddUnit(r)
	if data == false {
		http.Error(w, "Failed Insert Data", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "add-unit", http.StatusFound)
	}
}

func actionGetUnit(w http.ResponseWriter, r *http.Request) {
	data, err := er.GetUnit(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(result)
	}
}

func addExchangeView(w http.ResponseWriter, r *http.Request) {
	tmpl, data, err := er.AddExchangeView()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderHTML(w, tmpl, data)
}

func actionAddExchange(w http.ResponseWriter, r *http.Request) {
	data := er.AddExchange(r)
	if data == false {
		http.Error(w, "Failed Insert Data", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "add-exchange", http.StatusFound)
	}
}

func removeExchangeView(w http.ResponseWriter, r *http.Request) {
	tmpl, data, err := er.RemoveExchangeView()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderHTML(w, tmpl, data)
}

func actionRemoveExchange(w http.ResponseWriter, r *http.Request) {
	data := er.RemoveExchange(r)
	if data == false {
		http.Error(w, "Failed Delete Data", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "remove-exchange", http.StatusFound)
	}
}

func addExchangeRateView(w http.ResponseWriter, r *http.Request) {
	tmpl, data, err := er.AddExchangeRateView()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderHTML(w, tmpl, data)
}

func actionAddExchangeRate(w http.ResponseWriter, r *http.Request) {
	data := er.AddExchangeRate(r)
	if data == false {
		http.Error(w, "Failed Insert Data", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "add-exchange-rate", http.StatusFound)
	}
}

func viewExchangeRate(w http.ResponseWriter, r *http.Request) {
	tmpl, data, err := er.ViewExchangeRateView(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderHTML(w, tmpl, data)
}

func main() {
	db := er.InitDB()
	defer db.Close()

	/* Track Currency History */
	http.HandleFunc("/", viewExchangeRate)

	/* Currency Unit */
	http.HandleFunc("/add-unit", addUnitView)
	http.HandleFunc("/action-add-unit", actionAddUnit)

	/* Currency Exchange */
	http.HandleFunc("/add-exchange", addExchangeView)
	http.HandleFunc("/action-get-unit", actionGetUnit)
	http.HandleFunc("/action-add-exchange", actionAddExchange)
	http.HandleFunc("/delete-exchange", removeExchangeView)
	http.HandleFunc("/action-delete-exchange", actionRemoveExchange)

	/* Currency Exchange Rate */
	http.HandleFunc("/add-exchange-rate", addExchangeRateView)
	http.HandleFunc("/action-add-exchange-rate", actionAddExchangeRate)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
