package api

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)

		part := []string{"id, snippet"}
		call := yts.Videos.List(part).Chart("mostPopular").MaxResults(3)
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("error calling youtube api: %v", err)
		}

		fmt.Println("res", res)

		return c.JSON(fasthttp.StatusOK, res)
	}
}
