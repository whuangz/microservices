package mysql

import (
	"github.com/gocraft/dbr/v2"
	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
)

func Init(dataSource string) *dbr.Session {
	client := Connect(dataSource)
	return client
}

func Connect(dataSource string) *dbr.Session {
	conn, err := dbr.Open("mysql", dataSource, nil)

	if err != nil {
		logrus.Error(err)
	} else {
		session := conn.NewSession(nil)
		return session
	}
	return nil
}

// Transactional
const TxKey = "Tx"

func TransactionHandler(db *dbr.Session) echo.MiddlewareFunc {
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
