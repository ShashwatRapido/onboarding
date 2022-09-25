package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Review struct {
	PRODUCT_ID int
	Rating     int
	Review     string
}

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

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	p := []struct {
		ID          int
		Name        string
		Price       float32
		Description string
		Reviews     []Review
	}{
		{
			ID:          1,
			Name:        "Macbook Pro 16\"",
			Price:       160000.00,
			Description: "Macbook Pro 16\" with Retina Display, i7, 16GB 512 GB SSD.",
		},
		{
			ID:          2,
			Name:        "Dell 24\" IPS Monitor",
			Price:       14000.00,
			Description: "Dell Monitor with  IPS, USB Type C Connectivity and VESA Mount",
		},
	}

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, p)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		id_i, _ := strconv.Atoi(id)
		if id_i > len(p) || id_i < 1 {
			c.String(http.StatusNotFound, "Product does not exist")
			return
		}
		p := p[id_i-1]
		reviews, err := getReviews(id_i)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Something went wrong. Reason : %v", err))
		} else {
			p.Reviews = reviews
			c.JSON(http.StatusOK, p)
		}

	})

	router.Run(":9090")
}

func getReviews(id int) ([]Review, error) {
	u := fmt.Sprintf("%v/products/%v/reviews", os.Getenv("REVIEW_SVC_HOST"), id)
	var c = &http.Client{Timeout: 10 * time.Second}

	res, err := c.Get(u)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	reviews := &[]Review{}
	if res.StatusCode != 200 {
		return *reviews, err
	}
	err = json.NewDecoder(res.Body).Decode(reviews)
	return *reviews, err

}
