package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"insanitygaming.net/bans/src/gb/services/logger"
)

const TIME_OUT time.Duration = 5 * time.Second

type Database struct {
	db *sql.DB
}

func New() *Database {

	return &Database{}
}

func (db *Database) Connect() {
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
	db.db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	db.db.SetMaxOpenConns(10)
	db.db.SetMaxIdleConns(3)
	db.db.SetConnMaxLifetime(time.Minute * 5)
	logger.Logger().Info("Connected to database")
}

func (db *Database) Close() {
	db.db.Close()
	logger.Logger().Debug("Closed database connection")
}

//TODO: Extract setup to a separate app or something
//		I just dislike having a flag to toggle this just doesnt
//		seem like the correct place to do this

func (db *Database) RunSetup(ctx context.Context) {
	logger.Logger().Info("Running database setup")

	_, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_admin (
		admin_id BIGINT UNSIGNED NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		auths VARCHAR(255) DEFAULT '{}',
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL,
		adm_groups VARCHAR(32) NOT NULL DEFAULT '',
		web_groups VARCHAR(32) NOT NULL DEFAULT '',
		svr_groups VARCHAR(32) NOT NULL DEFAULT '',
		flags VARCHAR(32) NOT NULL DEFAULT '',
		immunity INT UNSIGNED NOT NULL DEFAULT 0
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_admin: ", err)
		return
	}
	logger.Logger().Info("Admins database was successfully created")

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_group (
		group_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(64) NOT NULL UNIQUE,
		flags INT NOT NULL DEFAULT 0,
		immunity INT UNSIGNED NOT NULL DEFAULT 0,
		group_type INT UNSIGNED DEFAULT 0
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_group: ", err)
		return
	}
	logger.Logger().Info("Groups database was successfully created")

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_server (
		server_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		app_id INT UNSIGNED NOT NULL,
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

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_server_group (
		group_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(64) NOT NULL UNIQUE,
		servers VARCHAR(32) NOT NULL DEFAULT ''
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_server_group: ", err)
		return
	}
	logger.Logger().Info("Server Groups database was successfully created")

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_bans (
		ban_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		admin_id BIGINT UNSIGNED NOT NULL,
		ban_tyoe INT UNSIGNED NOT NULL,
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
	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_application (
		application_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		image VARCHAR(255) NULL
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_application: ", err)
		return
	}
	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS gb_ban_type (
		ban_type_id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL
	);`)
	if err != nil {
		logger.Logger().Error("Error setting up gb_ban_type: ", err)
		return
	}
	logger.Logger().Info("Bans database was successfully created")
	logger.Logger().Info("Database setup was successful")
}

func (db *Database) QueryRow(background context.Context, query string, args ...interface{}) (*sql.Row, error) {
	ctx, cancel := context.WithTimeout(background, time.Second*TIME_OUT)
	row := db.db.QueryRowContext(ctx, query, args...)
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

func (db *Database) Query(background context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(background, time.Second*TIME_OUT)
	rows, err := db.db.QueryContext(ctx, query, args...)
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

func (db *Database) Exec(background context.Context, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(background, time.Second*TIME_OUT)
	result, err := db.db.ExecContext(ctx, query, args...)
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
