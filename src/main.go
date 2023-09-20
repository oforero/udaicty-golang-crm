package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"udacity_crm/backend"
	"udacity_crm/db"
	"udacity_crm/model"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/uptrace/bun"
)

type customer model.Customer
type customers []model.Customer

func (data customers) WriteResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(data)
}

func (data customer) WriteResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(data)
}

func serveFile(folder string, mime string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if file, ok := vars["file"]; !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {

			fqn := fmt.Sprintf("./%s/%s", folder, file)
			fmt.Printf("Serving file: %v\n", fqn)
			fi, err := os.Open(fqn)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			defer func() {
				if err := fi.Close(); err != nil {
					panic(err)
				}
			}()
			w.Header().Set("Content-Type", mime)
			w.WriteHeader(http.StatusOK)
			// make a buffer to keep chunks that are read
			buf := make([]byte, 1024)
			for {
				// read a chunk
				n, err := fi.Read(buf)
				if err != nil && err != io.EOF {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if n == 0 {
					break
				}

				// write a chunk
				if _, err := w.Write(buf[:n]); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	}
}

var css backend.Route = backend.Route{
	Url:     "/css/{file}",
	Handler: serveFile("css", "text/css"),
	Method:  "GET",
}

var scripts backend.Route = backend.Route{
	Url:     "/scripts/{file}",
	Handler: serveFile("scripts", "text/javascript"),
	Method:  "GET",
}

var swagger backend.Route = backend.Route{
	Url:     "/docs/{file}",
	Handler: serveFile("docs", "text/html"),
	Method:  "GET",
}

var root backend.Route = backend.Route{
	Url: "/",
	Handler: func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/index.html", http.StatusSeeOther)
	},
	Method: "GET",
}

var getCustomers backend.Route = backend.Route{
	Url:    "/customers",
	Method: "GET",
}

var getCustomerById backend.Route = backend.Route{
	Url:    "/customers/{ID}",
	Method: "GET",
}

var addCustomer backend.Route = backend.Route{
	Url:    "/customers",
	Method: "POST",
}

var updateCustomer backend.Route = backend.Route{
	Url:    "/customers/{ID}",
	Method: "PATCH",
}

var deleteCustomer backend.Route = backend.Route{
	Url:    "/customers/{ID}",
	Method: "DELETE",
}

func setupRoutesWithHandlers(dbServer *db.DB, ctx context.Context) *mux.Router {
	getCustomers.Handler = backend.BuildGetAllCustomersHandler(dbServer, ctx)
	getCustomerById.Handler = backend.BuildGetCustomerByIdHandler(dbServer, ctx)
	addCustomer.Handler = backend.BuildAddCustomerHandler(dbServer, ctx)
	updateCustomer.Handler = backend.BuildUpdateCustomerHandler(dbServer, ctx)
	deleteCustomer.Handler = backend.BuildDeleteCustomerHandler(dbServer, ctx)

	// addCustomer.Handler = backend.NewElemenbByIdHandler(fromDbAddCustomer)
	routes := []backend.Route{root, css, scripts, swagger, getCustomerById, addCustomer, updateCustomer, deleteCustomer}

	r := backend.BuildRouter(routes)

	return r
}

//	@title			Udacity Go Class CRM API
//	@version		1.0
//	@description	This is the CRM API for the Udacity Go Language Course

//	@contact.name	API Support
//	@contact.url	http://gocrm.io/support
//	@contact.email	support@gocrm.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @host						localhost:3000
// @BasePath					/
// @query.collection.format	multi
func main() {
	ctx := context.Background()
	dbServer := db.NewSqliteDB()
	bunDB := bun.DB(*dbServer)

	model.InitModel(&bunDB, ctx)
	model.InitData(&bunDB, ctx)

	router := setupRoutesWithHandlers(dbServer, ctx)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	http.ListenAndServe(":3000", corsOpts.Handler(router))
}
