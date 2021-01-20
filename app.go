// app.go

package main

import (
	"database/sql"
    "fmt"
    "log"
    "io"

    "net/http"
    "encoding/json"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, host, dbname string) {
    connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
    fmt.Printf("DB Connection %s\n", connectionString)
    var err error
    a.DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

func (a *App) Run(addr string) {
    fmt.Println("Server Listening at :9100")
    log.Fatal(http.ListenAndServe(":9100", a.Router))
 }

 func getStatus(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello At : 9100")
	io.WriteString(w, "It's OK!\n")
}

func (a *App) initializeRoutes() {
   // a.Router.HandleFunc("/products", a.createUsers).Methods("GET")
    a.Router.HandleFunc("/status", getStatus)
    a.Router.HandleFunc("/user", a.createUser).Methods("POST")
    a.Router.HandleFunc("/users", a.getAllUsers).Methods("GET")
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
    var u user
    decoder := json.NewDecoder(r.Body)
    fmt.Printf("Create User API, Request Body : %s \n", r.Body)
    if err := decoder.Decode(&u); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := u.createUser(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, u)
}

func (a *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Get Users API \n")
    users, err := getUsers(a.DB, 0, 0)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, users)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}