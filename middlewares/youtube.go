package middlewares

import (
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YoutubeService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := os.Getenv("API_KEY")
			ctx := context.Background()

			service, err := youtube.NewService(ctx, option.WithAPIKey(key))
			if err != nil {
				logrus.Fatalf("error creating youtube service: %v", err)
			}

			c.Set("yts", service)
			if err := next(c); err != nil {
				return err
			}
			return nil
		}
	}
}
