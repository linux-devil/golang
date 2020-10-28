package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/store"
)

var authenticator auth.Authenticator
var cache store.Cache

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/token", createToken).Methods("GET")
	router.HandleFunc("/v1/book/{id}", getBookAuthor).Methods("GET")
	log.Println("server started and listening on http://127.0.0.1:8082 !!")
	http.ListenAndServe("127.0.0.1:8082", router)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	body := fmt.Sprintf("token: %s \n", token)
	w.Write([]byte(body))
}

func getBookAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	books := map[string]string{
		"1449311601": "Harsht",
		"7777":       "spark",
	}
	body := fmt.Sprintf("Author: %s \n", books[id])
	w.Write([]byte(body))
}

func validateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {

	if userName == "medium" && password == "medium" {
		return auth.NewDefaultUser("medium", "1", nil, nil), nil
	}

	return nil, fmt.Errorf("invalid credentials")
}

func setupGoGuardian() {

	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*10)

	basicStrategy := basic.New(validateUser, cache)
	tokenStrategy := bearer.New(bearer.NoOpAuthenticate, cache)

	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)

}

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("executing Auth middleware")
		user, err := authenticator.Authenticate(r)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Printf("User %s Authenticate\n", user.UserName())
		next.ServeHTTP(w,r)
	}
}
