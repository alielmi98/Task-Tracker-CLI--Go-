# Task Tracker CLI - Go Implementation

This project implements a command-line interface (CLI) application for tracking tasks. It's a solution to the [Task Tracker project](https://roadmap.sh/projects/task-tracker) found on roadmap.sh.

This Go implementation adheres to the project's requirements, using only native Go modules and JSON file storage for task persistence.

## Features

* Add tasks: Add new tasks with descriptions.
* Update tasks: Modify existing task descriptions.
* Delete tasks: Remove tasks.
* Mark tasks as in progress/done: Change task status.
* List tasks: Display all tasks, or filter by status (`todo`, `in-progress`, `done`).

## Requirements

* Go 1.18 or higher installed on your system.

## Installation

1. Clone the repository:
   ```bash
   git clone <repository_url>

2. Navigate to the cmd directory:
    ```bash
    cd <project_directory>/cmd/task-tracker

3. Build the application:
    ```bash
    go build

4. Run the application:
    ```bash
    ./task-tracker <command> <arguments>

## Usage
The application is run from the command line. The following commands are supported:

* add <description>: Adds a new task with the given description.
* update <id> <new_description>: Updates the description of the task with the given ID.
* delete <id>: Deletes the task with the given ID.
* mark-in-progress <id>: Marks the task with the given ID as "in-progress".
* mark-done <id>: Marks the task with the given ID as "done".
* list: Lists all tasks.
* list <status>: Lists tasks with the specified status ("todo", "in-progress", or "done").

## Example Usage
```bash
    ./task-tracker add "Buy groceries"
    ./task-tracker update 1 "Buy groceries and milk"
    ./task-tracker list done
    ./task-tracker delete 1
```
The application will store tasks in a file named tasks.json in the <project_directory>/internal/data directory.

## Project Structure
```bash
task-tracker/
├── cmd/
│   └── task-tracker/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── data/
│   │   └── tasks.json
│   └── pkg/
│       └── ...
└── go.mod
```
## Design Patterns and Best Practices

This project utilizes several design patterns and Go programming best practices to structure and organize the code:

**1. Interface Segregation Principle (ISP):** The `Task` interface contains only the essential methods for each task, avoiding a large, bloated interface. This allows different implementations of the interface to only implement the methods they need, preventing unnecessary method implementations.

**2. Dependency Inversion Principle (DIP):** `TaskManagerImpl` depends on the `Task` interface, not on a concrete implementation of `Task`. This makes `TaskManagerImpl` flexible and allows it to work with any implementation of the `Task` interface without modification.

**3. Repository Pattern:** `TaskManagerImpl` acts as a repository. This layer is responsible for storing and retrieving data (Tasks) and is separated from the business logic. This separation improves testability and maintainability.

**4. Factory Pattern (Implicit):** The `NewTaskManager` function acts as a factory, creating instances of `TaskManagerImpl`. This centralizes the creation and configuration of `TaskManagerImpl`.

**5. Error Handling:** The code uses `error` effectively for error handling. Errors are reported clearly, and `fmt.Errorf` is used to create descriptive error messages.

**6. Use of `filepath.Abs`:**  `filepath.Abs` ensures the use of absolute paths for the task storage file, making the code portable across different systems.

**7. Appropriate Data Structures:** Using `[]Task` for storing tasks provides efficient access and performance.

**8. Code Cleanliness:** The code is well-formatted and includes clear comments, improving readability and understanding.

## Error Handling
The application includes basic error handling to manage file I/O and data parsing issues. More robust error handling could be added for production use.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
MIT License

## Roadmap.sh Project Page
This project is a solution for the Task Tracker project idea found on roadmap.sh. This README file provides instructions for running the Go implementation of this project.