package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"insanitygaming.net/bans/services/logger"
)

var db *sql.DB

const TIME_OUT time.Duration = 5 * time.Second

func Connect() {
	config := mysql.NewConfig()
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Net = "tcp"
	config.Addr = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_NAME")
	config.Loc = time.Local
	config.ParseTime = true
	var err error
	fmt.Println(config.FormatDSN())
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(time.Minute * 5)
	logger.Logger().Info("Connected to database")
}

func Close() {
	db.Close()
	logger.Logger().Debug("Closed database connection")
}

func RunSetup() {
	logger.Logger().Info("Running database setup")

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS gb_admins (
		admin_id BIGINT UNSIGNED NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		groups VARCHAR(32) NOT NULL DEFAULT '',
		flags VARCHAR(32) NOT NULL DEFAULT '',
		immunity INT UNSIGNED NOT NULL DEFAULT 0
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_admins: ", err)
		return
	}
	logger.Logger().Info("Admins database was successfully created")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gb_groups (
		group_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(64) NOT NULL UNIQUE,
		flags INT NOT NULL DEFAULT 0,
		group_type INT UNSIGNED NOT NULL,
		immunity INT UNSIGNED NOT NULL DEFAULT 0
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_groups: ", err)
		return
	}
	logger.Logger().Info("Groups database was successfully created")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gb_servers (
		server_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		ip VARCHAR(64) NOT NULL,
		port INT UNSIGNED NOT NULL,
		immunity INT UNSIGNED NOT NULL DEFAULT 0,
		UNIQUE KEY ip_port (ip, port)
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_servers: ", err)
		return
	}
	logger.Logger().Info("Servers database was successfully created")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gb_server_groups (
		group_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(64) NOT NULL UNIQUE,
		servers VARCHAR(32) NOT NULL DEFAULT ''
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_server_groups: ", err)
		return
	}
	logger.Logger().Info("Server Groups database was successfully created")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gb_bans (
		ban_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		admin_id BIGINT UNSIGNED NOT NULL,
		server_id INT UNSIGNED NOT NULL,
		player_id BIGINT UNSIGNED NOT NULL,
		reason VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL,
		expires_at DATETIME NOT NULL,
		comment VARCHAR(255) NULL,
		extra LONGTEXT NULL
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_bans: ", err)
		return
	}
	//TODO: Add in application sql
	logger.Logger().Info("Bans database was successfully created")
	logger.Logger().Info("Database setup was successful")
}

func QueryRow(background context.Context, query string, args ...interface{}) (*sql.Row, error) {
	ctx, cancel := context.WithTimeout(background, time.Second*TIME_OUT)
	row := db.QueryRowContext(ctx, query, args...)
	defer func() {
		logger.Logger().Warnf("Query took too long to execute\nQuery:%s\nArgs: %v", query, args)
		cancel()
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("Query took too long to execute")
	default:
		if row.Err() != nil {
			logger.Logger().Error("Error executing query: " + row.Err().Error())
			return nil, row.Err()
		}
		return row, nil
	}
}

func Query(background context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(background, time.Second*TIME_OUT)
	rows, err := db.QueryContext(ctx, query, args...)
	defer func() {
		logger.Logger().Warnf("Query took too long to execute\nQuery:%s\nArgs: %v", query, args)
		cancel()
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("Query took too long to execute")
	default:
		if err != nil {
			logger.Logger().Error("Error executing query: " + err.Error())
			return nil, err
		}
		return rows, nil
	}
}

func Exec(background context.Context, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(background, time.Second*TIME_OUT)
	result, err := db.ExecContext(ctx, query, args...)
	defer func() {
		logger.Logger().Warnf("Query took too long to execute\nQuery:%s\nArgs: %v", query, args)
		cancel()
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("Query took too long to execute")
	default:
		if err != nil {
			logger.Logger().Error("Error executing query: " + err.Error())
			return nil, err
		}
		return result, nil
	}
}
