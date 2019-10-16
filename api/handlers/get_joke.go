package handlers

import (
	"fmt"
	"net/http"

	"github.com/random_name_svc/models"
	"github.com/random_name_svc/service"
)

func getRandomJokeByName(w http.ResponseWriter, r *http.Request) {
	resp, err := service.GetRandomJokeByName()
	if err != nil {
		errMsg := models.MarshalErrorBinary(&models.Errors{ErrorCode: http.StatusInternalServerError, Msg: err.Error()})
		fmt.Fprintln(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(models.MarshalResponseBinary(&models.Response{Joke: resp}))
}
