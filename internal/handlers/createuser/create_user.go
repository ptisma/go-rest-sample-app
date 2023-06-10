package createuser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ptisma/go-rest-sample-app/cmd/api/models"
	"github.com/ptisma/go-rest-sample-app/pkg/application"
	"github.com/ptisma/go-rest-sample-app/pkg/middleware"
)

func createUser(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		user := &models.User{}
		json.NewDecoder(r.Body).Decode(user)

		if err := user.Create(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	return middleware.Chain(createUser(app), middleware.LogRequest)
}
