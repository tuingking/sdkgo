package osql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

const (
	defaultMaxConnectTimeout = 15 * time.Second
)

type OSql interface {
	DB() *sqlx.DB
	Stop() error
}

type osql struct {
	opt Option
	db  *sqlx.DB
}

type Option struct {
	Enabled    bool
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	Db         string
	Connection ConnectionOpt
}

type ConnectionOpt struct {
	MaxOpen     int
	MaxIdle     int
	MaxLifeTime time.Duration
}

func Init(opt Option) OSql {
	if !opt.Enabled {
		log.Printf(`[osql] db "%s/%s" is not enabled`, opt.Host, opt.Db)
		return nil
	}

	osql := &osql{
		opt: opt,
	}
	if err := osql.initDB(); err != nil {
		log.Fatalf("failed init db. err: %v", err)
	}

	return osql
}

func (osql *osql) DB() *sqlx.DB {
	return osql.db
}

func (osql *osql) Stop() error {
	return osql.db.Close()
}

func (o *osql) initDB() error {
	connectionString := o.connectionString()
	db, err := sql.Open(o.opt.DriverName, connectionString)
	if err != nil {
		log.Fatalf("[osql] cannot connect to database. err: %v", err)
	}

	sqlxDB := sqlx.NewDb(db, o.opt.DriverName)
	sqlxDB.SetMaxOpenConns(o.opt.Connection.MaxOpen)
	sqlxDB.SetMaxIdleConns(o.opt.Connection.MaxIdle)
	sqlxDB.SetConnMaxLifetime(o.opt.Connection.MaxLifeTime)

	// retry mechanism if couldn't connect to database
	sqlOpen := func() error {
		log.Printf(`[osql] ping "%s:%s"...`, o.opt.Host, o.opt.Port)
		return db.PingContext(context.Background())
	}
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = defaultMaxConnectTimeout
	bo.MaxInterval = defaultMaxConnectTimeout
	bo.Multiplier = backoff.DefaultMultiplier
	bo.RandomizationFactor = backoff.DefaultRandomizationFactor
	err = backoff.RetryNotify(
		sqlOpen,
		bo,
		backoff.Notify(func(err error, duration time.Duration) {
			if err != nil {
				log.Printf("[osql] retry. err=%v (%d)", err, duration)
			}
		}))
	if err != nil {
		return errors.Wrap(err, "retry failed")
	}

	o.db = sqlxDB
	log.Printf("[osql] successfully connected to %s\n", o.opt.Db)
	return nil
}

func (o *osql) connectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", o.opt.User, o.opt.Password, o.opt.Host, o.opt.Port, o.opt.Db)
}
