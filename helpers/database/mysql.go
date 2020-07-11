package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

func Init(dataSource string) *sql.DB {
	client := Connect(dataSource)
	return client
}

func Connect(dataSource string) *sql.DB {
	conn, err := sql.Open("mysql", dataSource)

	if err != nil {
		logrus.Error(err)
	} else {
		return conn
	}
	return nil
}

// Transactional
const TxKey = "Tx"

func TransactionHandler(db *sql.DB) echo.MiddlewareFunc {
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
