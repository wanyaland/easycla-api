// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"regexp"
	"time"

	"github.com/communitybridge/easycla-api/config"
	log "github.com/communitybridge/easycla-api/logging"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func initDB(config config.Config) *sqlx.DB {
	var d *sqlx.DB
	if false {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=verify-full",
			config.RDSUsername, config.RDSPassword, config.RDSHost, config.RDSPort, config.RDSDatabase)
		log.Infof("Initializing DB connection to database: %s", connStr)
		var err error
		d, err = sqlx.Connect("postgres", connStr)
		if err != nil {
			log.Panicf("unable to connect to database: %s on host: %s:%d with user: %s. Error: %v",
				config.RDSDatabase, config.RDSHost, config.RDSPort, config.RDSUsername, err)
		}
	}

	if true {
		dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require connect_timeout=5",
			config.RDSHost, config.RDSPort, config.RDSUsername, config.RDSPassword, config.RDSDatabase)
		log.Infof("Initializing DB connection to database: %s", stripPassword(dbInfo))
		var err error
		d, err = sqlx.Open("postgres", dbInfo)
		if err != nil {
			log.Panicf("unable to connect to database: %s on host: %s with user: %s. Error: %v",
				config.RDSDatabase, config.RDSHost, config.RDSUsername, err)
		}
	}

	d.SetMaxOpenConns(viper.GetInt("DB_MAX_CONNECTIONS"))
	d.SetMaxIdleConns(5)
	d.SetConnMaxLifetime(15 * time.Minute)

	return d
}

func stripPassword(dbConnection string) string {
	re := regexp.MustCompile(`password=\w+\s+`)
	return re.ReplaceAllString(dbConnection, "password=<redacted> ")
}
