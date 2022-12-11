package middlewares

import (
	"youtube-manager/databases"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type DatabaseClient struct {
	DB *gorm.DB
}

func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db, _ := databases.Connect()
			d := DatabaseClient{DB: db}
			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}
