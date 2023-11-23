package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samueldsiqueira/go-initial-learn/internal/infra/database"
	"github.com/samueldsiqueira/go-initial-learn/internal/usecase"
)

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() { //quero essa função referente a structure CAR, fazendo relacionamento = metodo
	println(c.Model + " has been started")
}

func (c Car) ChangeColor(color string) {
	c.Color = color //duplica o valor na memoria dentro desse escopo, mas para alterar o original, deve-se orientar o GO informando o ponteiro por isso deve-se adicionar um *(asterico) informando o ponteiro
	println("New color: " + c.Color)
}

// função comum
func soma(x, y int) int {
	return x + y
}

// func main() {
// 	car := Car{ //declarando e atribuindo uma variavel
// 		Model: "Ferrari",
// 		Color: "Red",
// 	}
// 	car.Start()
// 	car.ChangeColor("Blue")
// 	println(car.Color)
// 	// println(car.Model)
// }

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
