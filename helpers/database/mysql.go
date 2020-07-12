package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

func Init(dataSource string) *sqlx.DB {
	client := Connect(dataSource)
	return client
}

func Connect(dataSource string) *sqlx.DB {
	conn, err := sqlx.Connect("mysql", dataSource)

	if err != nil {
		logrus.Error(err)
	} else {
		return conn
	}
	return nil
}

// Transactional
const TxKey = "Tx"

func TransactionHandler(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tx, _ := db.Begin()
			c.Set(TxKey, tx)
			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debug("Transction Rollback: ", err)
				return err
			}
			logrus.Debug("Transaction Commit")
			tx.Commit()
			return nil
		}
	}
}
