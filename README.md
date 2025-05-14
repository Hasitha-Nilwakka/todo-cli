# todo-cli

A simple command-line tool to manage your tasks.
The tool stores the tasks in JSON format in the project root.

## Usage

```sh
./todo [add|list|complete|delete] [arguments]
```

# Commands

## add "task description"
Adds a new task with the given description.

Example:
```sh
./todo add "Buy groceries"
```
## list
Lists all tasks.

Example:
```sh
./todo list
```

## complete "task number"
Marks the specified task as completed.

Example:
```sh
./todo complete 1
```

## delete "task number"
Deletes the specified task.

Example:
```sh
./todo delete 1
```

# To launch this application

```sh
git clone https://github.com/Hasitha-Nilwakka/todo-cli.git
cd todo-cli
```
