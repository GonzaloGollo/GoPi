package main

import (
	//transformamos a json
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GonzaloGollo/GoPi/internal/domain"
	"github.com/GonzaloGollo/GoPi/internal/user"
)

func main() {
	server := http.NewServeMux()

	dbCmd := user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "Lelo",
			LastName:  "Pepe",
			Email:     "pepe@gmail.com",
		}, {
			ID:        2,
			FirstName: "Lelo2",
			LastName:  "Pepe2",
			Email:     "pepe2@gmail.com",
		}, {
			ID:        3,
			FirstName: "Lelo3",
			LastName:  "Pepe3",
			Email:     "pepe3@gmail.com",
		}},
		MaxUserID: 3,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	repo := user.NewRepo(dbCmd, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background() //ctx para pasarle info de fondo a las diferentes capas

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
