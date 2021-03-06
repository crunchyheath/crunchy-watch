package util

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const DEFAULT_HEALTHCHECK_QUERY = "SELECT 1;"

func HealthCheck(target string) error {
	// Create connection to the datasource
	conn, err := sql.Open("postgres", target)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	defer conn.Close()

	// Check that the connection is alive
	err = conn.Ping()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// Check that the connection can be queried
	rows, err := conn.Query(DEFAULT_HEALTHCHECK_QUERY)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	defer rows.Close()

	return nil
}
