package _handler

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func dsn(dbName string) string {
	username := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	hostname := os.Getenv("PG_HOST")
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, password, hostname, dbName)
}

func Connect2PGsql(dbName string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dsn(dbName))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func ExecQuery(query string, dbName string) error {
	conn := Connect2PGsql(dbName)
	defer conn.Close(context.Background())
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return err
}

/*
func InsertLSVData(lsvData _models.LSVData, conn *pgx.Conn) ([]_models.LSVData, error) {

	if _, err := conn.Exec(context.Background(), "INSERT INTO USERS(USERNAME) VALUES($1)", u.UserName); err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return
	}
	fmt.Println("Insertion Succesfull")
}

func LSVPlanningData(conn *pgx.Conn) ([]_models.LSVData, error) {
	var row _models.LSVData
	var rowdata []_models.LSVData
	query := "Select id, prime, primehal, marketorplant, fgxreporting, eoy, status, enddate, fmm, tmm, vm, application, appname, cmdbid, appstatus, migrated, dept, consumption, created from cloudcosts"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var id int
	var prime string
	var primehal string
	var marketorplant string
	var fgxreporting string
	var eoy int
	var status string
	var enddate string
	var fmm string
	var tmm string
	var vm string
	var application string
	var appname string
	var cmdbid string
	var appstatus string
	var migrated string
	var dept string
	var consumption string
	var created string

	for rows.Next() {

		err = rows.Scan(&id, &prime, &primehal, &marketorplant, &fgxreporting, &eoy, &status,
			&enddate, &fmm, &tmm, &vm, &application, &appname, &cmdbid, &appstatus, &migrated, &dept, &consumption, &created)
		if err != nil {
			return nil, err
		}

		row.Id = id
		row.Prime = prime
		row.PrimeHal = primehal
		row.MarketOrPlant = marketorplant
		row.FGXReporting = fgxreporting
		row.EOY = eoy
		row.Status = status
		row.EndDate = enddate
		row.FMM = fmm
		row.TMM = tmm
		row.VM = vm
		row.Application = application
		row.AppName = appname
		row.CMDBID = cmdbid
		row.AppStatus = appstatus
		row.Migrated = migrated
		row.Dept = dept
		row.Consumption = consumption
		row.Created = created
		rowdata = append(rowdata, row)
	}

	return rowdata, err
}
*/
