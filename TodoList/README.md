# TodoList Written in Go

# NOTE WINDOWS VERSION IS NOT WORKING.



[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)


## Clone the project

```
git clone https://github.com/OkRespire/go-projects.git
cd go-projects/TodoList
```
---
## Build the project & initialise it
```
go build
.\respireToDoList.exe init  //if on windows
./respireToDoList init //if on linux
```
---
## Run the project

```
.\respireToDoList.exe init  //if on windows
./respireToDoList init //if on linux
```


---
## Usage
| Default | Left align |
| :- | :- | 
| init | adds the required files needed to use the to do list|
| add | appends the required task to the CSV file | 
| delete | deletes an entry in the CSV file  | 
| complete | Sets the "completion column" to TRUE. |
| list |Lists all uncompleted tasks. If you want to see all the tasks, uses the -a flag.  |