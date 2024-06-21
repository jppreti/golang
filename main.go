package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Cliente struct {
	ID    int
	Nome  string
	Email string
}

func main() {
	cliente := Cliente{
		ID:    1,
		Nome:  "João",
		Email: "joao@gmail.com",
	}
	fmt.Println(cliente.ID)
	err := SaveCliente(cliente)
	if err != nil {
		panic(err)
	}

	//http.HandleFunc("/", HomeHandler)
	//http.ListenAndServe(":5000", nil)
	// Substituir pelo framework echo
	e := echo.New()
	e.POST("/clientes", CreateCliente)
	e.Logger.Fatal(e.Start(":5000"))
	println("Olá Mundo!!!")
}

func CreateCliente(c echo.Context) error {
	cliente := Cliente{}
	if err := c.Bind(&cliente); err != nil {
		return err
	}
	err := SaveCliente(cliente)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, cliente)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!!!"))
}

func SaveCliente(cliente Cliente) error {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cliente (id, nome, email) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cliente.ID, cliente.Nome, cliente.Email) //_ ignorar o resultado que é o número de linhas inseridas
	if err != nil {
		return err
	}
	return nil
}
