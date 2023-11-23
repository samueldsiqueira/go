package main

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

func main() {
	car := Car{ //declarando e atribuindo uma variavel
		Model: "Ferrari",
		Color: "Red",
	}
	car.Start()
	car.ChangeColor("Blue")
	println(car.Color)
	// println(car.Model)
}
