package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justtaldevelops/go-hcaptcha"
)

type TOKEN struct {
	URL     string `json:"url" binding:"required"`
	API_KEY string `json:"api_key" binding:"required"`
}

func getToken(con *gin.Context) {
	var data TOKEN
	con.BindJSON(&data)

	c, err := hcaptcha.NewChallenge(
		data.URL,
		data.API_KEY,
		hcaptcha.ChallengeOptions{
			Timeout: 10 * time.Second,
		},
	)
	if err != nil {
		con.IndentedJSON(
			http.StatusInternalServerError, gin.H{
				"error": "Ocorreu um erro, tente novamente!",
			})
	}
	err = c.Solve(&hcaptcha.GuessSolver{})
	if err != nil {
		con.IndentedJSON(
			http.StatusInternalServerError, gin.H{
				"error": "Erro solving, tente novamente!",
			})
	} else {
		con.JSON(
			http.StatusOK, gin.H{
				"token": c.Token(),
			})
	}
}

func get_port() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	router := gin.Default()
	router.POST("/token", getToken)
	router.Run(get_port())
}
