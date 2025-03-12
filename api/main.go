package main

import (
	"fmt"
	"net/http"

	"os"

	db "helm/app/test/db"
	"helm/app/test/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	type Test struct {
		Name string
	}
	ran_env := os.Getenv("TEST")
	_ = db.DbConnect()
	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": ran_env,
		})
	})

	r.Use(middleware.AuthorizationMiddleware())
	{
		r.GET("/api/data", func(c *gin.Context) {
			d := db.DbConnect()
			res, err := d.Query("SELECT * FROM test;")
			if err != nil {
				fmt.Printf("Bad Query: %+v\n", err)
				c.JSON(http.StatusInternalServerError, "executed query was bad")
				return
			}
			var test Test
			for res.Next() {
				res.Scan(&test.Name)
			}
			c.JSON(http.StatusOK, test)
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
