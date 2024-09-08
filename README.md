# CLI todo list

CLI application for task management, using the [Go](https://go.dev/) language and the [Table](https://github.com/aquasecurity/table) library.

This project was made based on the video from the channel [Coding with Patrik](https://www.youtube.com/@codingwithpatrik).

## How to run on your machine
1. Clone the repository:
   ```
   git clone https://github.com/alexsandro49/cli_todo_list.git
   ```
2. Access the project files:
   ```
   cd cli_todo_list/cli_todo_list
   ```
3. Compile the application:
   ```
   go build .
   ```
4. Change the permissions of the binary:
   ```
   chmod +x ./todo
   ```
5. Run the application:
   ```
   ./todo --list
   ```

## How the application works
* `--help` - To list all commands.
* `--list` - Add a new todo specify title.
* `--add` - Add a new todo specify title.
* `--edit` - Edit a todo by index & specify a new title. id:new_title.
* `--delete` - Specify a todo by index to delete.
* `--toggle` - Specify a todo by index to toggle.

## References
### Original project video:
   ```
   https://youtu.be/g16Zf0KQEWI?si=eA7hnb8BxjFIJ1EB
   ```
### Original project repository:
   ```
   https://github.com/patni1992/CLI-Todo-App-In-Go
   ```
### Detailed article by the original author of the project
   ```
   https://codingwithpatrik.dev/posts/how-to-build-a-cli-todo-app-in-go
   ```

## License
- [MIT](https://github.com/alexsandro49/cli_todo_list/blob/main/LICENSE)
