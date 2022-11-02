package controllers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	reqApi "servers/restAPI/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var mapResponse map[string]interface{}

type Success struct {
	status  int
	message string
}

func ReadStatus200() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := make(chan *http.Response)
		errGrp, _ := errgroup.WithContext(context.Background())

		errGrp.Go(func() error {
			return reqApi.SendPostAsync("https://httpstat.us/200", []byte(""), result)
		})

		resultResponse := <-result
		defer resultResponse.Body.Close()

		bytes, _ := io.ReadAll(resultResponse.Body)

		json.Unmarshal([]byte(string(bytes)), &mapResponse)

		if resultResponse.StatusCode != 200 && mapResponse["message"] != "" {
			message := mapResponse["message"]
			c.JSON(resultResponse.StatusCode, gin.H{
				"message": message,
			})
			return
		}

		var responseSuccess Success
		responseSuccess.message = "success"
		responseSuccess.status = 200

		c.JSON(http.StatusOK, gin.H{
			"status":  responseSuccess.status,
			"message": responseSuccess.message,
		})
	}

}
