package todo

import (
	helper "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/utils"
	"net/http"
)


// InsertToDoHandler GET	/todos	fetches all todos
func FetchToDoAll(w http.ResponseWriter, r *http.Request) {

	if todos, err := helper.FetchAll(); err != nil {

		body := make(map[string]interface{})

		body["error"] = err

		utils.Respond(w, http.StatusInternalServerError, false, body)

	} else {

		body := make(map[string]interface{})

		body["todos"] = todos
		body["count"] = len(todos)

		utils.Respond(w, http.StatusOK, true, body)

	}

}