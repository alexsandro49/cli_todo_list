package main

func main() {
	todos := Todos{}

	storage := newStorage[Todos]("todos.json")
	storage.load(&todos)

	commandFlags := newCommandFlags()
	commandFlags.execute(&todos)

	storage.save(todos)
}
