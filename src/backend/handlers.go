package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ResultData interface {
	WriteResponse(http.ResponseWriter)
}

type allData func() ResultData
type elementById func(id int64) ResultData
type insertElement func[E]() ResultData

func NewListDataHandler(queryFn allData) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Searching the database with (NewListDataHandler): %v", r)
		data := queryFn()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data.WriteResponse(w)
	}
}

func NewElemenbByIdHandler(queryFn elementById) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Searching the database with (NewElemenbByIdHandler): %v\n", r)
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["ID"], 0, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
		} else {
			data := queryFn(id)
			fmt.Printf("Result: %v\n", data)
			if data == nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write(nil)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				json.NewEncoder(w).Encode(data)
			}

		}

	}
}
