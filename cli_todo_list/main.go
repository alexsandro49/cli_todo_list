package main

func main() {
	todos := Todos{}

	storage := newStorage[Todos]("todos.json")
	storage.load(&todos)

	todos.print()
	storage.save(todos)
}
