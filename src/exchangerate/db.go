package exchangerate

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "shopee-test"
)

var (
	db *sql.DB

	stmtInsertUnit, stmtGetAllUnit, stmtGetExchangeHistory     *sql.Stmt
	stmtInsertExchange, stmtRemoveExchange, stmtGetUnitExclude *sql.Stmt
	stmtGetAllExchange, stmtInsertExchangeRate                 *sql.Stmt
)

func InitDB() *sql.DB {
	connectDB()
	prepareStmt()

	return db
}

func connectDB() {
	postgres := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	var err error = nil

	db, err = sql.Open("postgres", postgres)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func prepareStmt() {
	var err error = nil

	stmtInsertUnit, err = db.Prepare("INSERT INTO currency_unit(unit, description) VALUES($1, $2)")
	if err != nil {
		panic(err)
	}
	stmtGetAllUnit, err = db.Prepare("SELECT id_unit, unit FROM currency_unit")
	if err != nil {
		panic(err)
	}
	stmtGetUnitExclude, err = db.Prepare("SELECT id_unit, unit FROM currency_unit WHERE id_unit != $1")
	if err != nil {
		panic(err)
	}
	stmtInsertExchange, err = db.Prepare("INSERT INTO currency_exchange(id_unit, id_unit_target, status) VALUES($1, $2, $3)")
	if err != nil {
		panic(err)
	}
	stmtRemoveExchange, err = db.Prepare("UPDATE currency_exchange SET status=0 WHERE id_exchange = $1")
	if err != nil {
		panic(err)
	}
	stmtGetAllExchange, err = db.Prepare(`
		SELECT ce.id_exchange, CONCAT_WS(' - ', uf.unit, ut.unit) AS unit
		FROM currency_exchange AS ce
		JOIN currency_unit AS uf ON ce.id_unit=uf.id_unit
		JOIN currency_unit AS ut ON ce.id_unit_target=ut.id_unit
		WHERE ce.status='1'
	`)
	if err != nil {
		panic(err)
	}
	stmtInsertExchangeRate, err = db.Prepare("INSERT INTO currency(date, id_exchange, value) VALUES($1, $2, $3)")
	if err != nil {
		panic(err)
	}
	stmtGetExchangeHistory, err = db.Prepare(`
		SELECT uf.unit AS unit_from, ut.unit AS unit_target, COALESCE(cr.value, 0),
			(SELECT COALESCE(AVG(c.value), 0)
			FROM currency c
			WHERE c.id_exchange=cr.id_exchange
			AND c.date > TO_DATE($1, 'yyyy-mm-dd') - INTERVAL '7 days') AS avg_rate
		FROM currency_exchange AS ce
		LEFT JOIN currency AS cr ON ce.id_exchange=cr.id_exchange
		LEFT JOIN currency_unit AS uf ON ce.id_unit=uf.id_unit
		LEFT JOIN currency_unit AS ut ON ce.id_unit_target=ut.id_unit
		ORDER BY uf.unit ASC
	`)
	if err != nil {
		panic(err)
	}
}

func insertUnit(unit string, description string) (bool, error) {
	_, err := stmtInsertUnit.Exec(unit, description)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getUnitExclude(id_unit string) ([]Currency, error) {
	res, err := stmtGetUnitExclude.Query(id_unit)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var (
		id   int
		unit string
	)

	data := []Currency{}
	for res.Next() {
		if err := res.Scan(&id, &unit); err != nil {
			return nil, err
		}

		data = append(data, Currency{
			ID:   id,
			Unit: unit,
		})
	}

	return data, nil
}

func getAllUnit() ([]Currency, error) {
	res, err := stmtGetAllUnit.Query()
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var (
		id   int
		unit string
	)

	data := []Currency{}
	for res.Next() {
		if err := res.Scan(&id, &unit); err != nil {
			return nil, err
		}

		data = append(data, Currency{
			ID:   id,
			Unit: unit,
		})
	}

	return data, nil
}

func insertExchange(id_unit string, id_unit_target string) (bool, error) {
	_, err := stmtInsertExchange.Exec(id_unit, id_unit_target, "1")
	if err != nil {
		return false, err
	}

	return true, nil
}

func removeExchange(id_exchange string) (bool, error) {
	_, err := stmtRemoveExchange.Exec(id_exchange)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getAllExchange() ([]CurrencyExchange, error) {
	res, err := stmtGetAllExchange.Query()
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var (
		id_exchange   int
		unit_exchange string
	)

	data := []CurrencyExchange{}
	for res.Next() {
		err := res.Scan(&id_exchange, &unit_exchange)
		if err != nil {
			return nil, err
		}

		data = append(data, CurrencyExchange{
			ID:   id_exchange,
			Unit: unit_exchange,
		})
	}

	return data, nil
}

func insertExchangeRate(date string, idExchangeRate string, rate string) (bool, error) {
	_, err := stmtInsertExchangeRate.Exec(date, idExchangeRate, rate)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getExchangeHistory(date string) ([]CurrencyExchangeRate, error) {
	res, err := stmtGetExchangeHistory.Query(date)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var (
		unitFrom   string
		unitTarget string
		rate       float64
		rateAvg    float64
	)

	data := []CurrencyExchangeRate{}
	for res.Next() {
		err := res.Scan(&unitFrom, &unitTarget, &rate, &rateAvg)
		if err != nil {
			return nil, err
		}

		data = append(data, CurrencyExchangeRate{
			UnitFrom:   unitFrom,
			UnitTarget: unitTarget,
			Rate:       rate,
			RateAvg:    rateAvg,
		})
	}

	return data, nil
}
