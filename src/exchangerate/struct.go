package exchangerate

type Currency struct {
	ID          int    `json:"id"`
	Unit        string `json:"unit"`
	Description string `json:"description,omitempty"`
}

type CurrencyExchange struct {
	ID   int    `json:"id"`
	Unit string `json:"unit"`
}

type CurrencyExchangeRate struct {
	UnitFrom   string  `json:"unit_from"`
	UnitTarget string  `json:"unit_target"`
	Rate       float64 `json:"rate"`
	RateAvg    float64 `json:"rate_avg"`
}
