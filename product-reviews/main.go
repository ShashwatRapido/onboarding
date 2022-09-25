package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" \"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	router.Use(gin.Recovery())

	type Review struct {
		PRODUCT_ID int
		Rating     int
		Review     string
	}
	reviews := make(map[int][]Review)
	reviews[1] = []Review{
		{
			PRODUCT_ID: 1,
			Rating:     5,
			Review:     "Excellent",
		},
		{
			PRODUCT_ID: 1,
			Rating:     4,
			Review:     "Good performance",
		},
	}

	reviews[2] = []Review{
		{
			PRODUCT_ID: 2,
			Rating:     3,
			Review:     "Needs 4K Upgrade",
		},
		{
			PRODUCT_ID: 2,
			Rating:     5,
			Review:     "have type c support",
		},
	}

	router.GET("/products/:id/reviews", func(c *gin.Context) {
		id := c.Param("id")
		id_i, _ := strconv.Atoi(id)

		if r, ok := reviews[id_i]; ok {
			c.JSON(http.StatusOK, r)
		} else {
			c.String(http.StatusNotFound, fmt.Sprintf("No reviews found for product %v", id))
		}

	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
