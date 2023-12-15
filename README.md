# Go Exercise

## Quick Start

### Build Docker Image and Run App in Container

```bash
docker build -t go-exercise .

docker run -p 8080:8080 go-exercise
```

### Run Unit Tests

```bash
go test ./... -v -coverprofile .testCoverage.txt
go tool cover -func .testCoverage.txt
```

Result:

```text
go-exercise/api/api.go:7:               RegisterRoutes  100.0%
go-exercise/api/task.go:31:             newTaskRoute    100.0%
go-exercise/api/task.go:40:             NewtaskHandler  100.0%
go-exercise/api/task.go:47:             GetTaskList     75.0%
go-exercise/api/task.go:66:             CreateTask      60.0%
go-exercise/api/task.go:88:             UpdateTask      50.0%
go-exercise/api/task.go:127:            DeleteTask      50.0%
go-exercise/api/task.go:149:            boolToInt       100.0%
go-exercise/api/task.go:156:            intToBool       100.0%
go-exercise/internal/restful.go:27:     ResponseSuccess 100.0%
go-exercise/internal/restful.go:42:     ResponseFail    100.0%
go-exercise/internal/restful.go:50:     NewErrorMessage 100.0%
go-exercise/service/task.go:33:         NewTaskService  100.0%
go-exercise/service/task.go:38:         CreateTask      83.3%
go-exercise/service/task.go:53:         GetTasks        100.0%
go-exercise/service/task.go:58:         UpdateTask      75.0%
go-exercise/service/task.go:71:         DeleteTask      100.0%
total:                                  (statements)    71.8%
```

### CRUD Examples

#### 1.  POST /task  (create task)

```json
request body
{
  "name": "買晚餐"
}

response status code 200
{
    "result": {"name": "買晚餐", "status": 0, "id": 1}
}
```

command memo:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "買晚餐"}' http://localhost:8080/task
```

#### 2.  GET /tasks (list tasks)

```json
{
    "result": [
        {"id": 1, "name": "買晚餐", "status": 0}
    ]
}
```

command memo:

```bash
curl http://localhost:8080/tasks
```

#### 3. PUT /task/<id> (update task)

```json
request body
{
  "name": "買早餐",
  "status": 1,
  "id": 1
}

response status code 200
{
  "result": {
    "name": "買早餐",
    "status": 1,
    "id": 1
  }
}
```

command memo:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "買早餐", "status": 1}' http://localhost:8080/task/1
```

#### 4. DELETE /task/<id> (delete task)

response status code 200

command memo:

```bash
curl -X DELETE http://localhost:8080/task/1
```
