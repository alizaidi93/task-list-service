# Task-List-Service

Task-List-Service is a simple Go service utilising hexagonal architecture design with Gorilla Mux & Postgres.

## Table of Contents  

- [Database table structure](#database-table-structure)
- [Running the service](#running-the-service)
- [Return task by UUID](#return-task-by-uuid)
  * [Request](#request)
  * [Response](#response)
- [Return all tasks created by user](#return-all-tasks-created-by-user)
  * [Request](#request-1)
  * [Response](#response-1)
- [Creates task](#creates-task)
  * [Request](#request-2)
  * [Response](#response-2)
- [Update a task](#update-a-task)
  * [Request](#request-3)
  * [Response](#response-3)

## Database table structure

> To run a postgres instance through docker : docker run  --name pg-docker -e POSTGRES_PASSWORD=<yourpassword> -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres

[Here is a  good article to run up postgres locally through a docker container.](https://hackernoon.com/dont-install-postgres-docker-pull-postgres-bee20e200198)

```sql
CREATE TABLE "public"."todo_list" (
    "id" SERIAL NOT NULL,
    "uuid" varchar NOT NULL,
    "task" varchar(10000) NOT NULL,
    "completed" bool NOT NULL,
    "createdby" varchar(100) NOT NULL,
    "modifiedby" varchar(100) NOT NULL,
    "modifiedon" timestamp NOT NULL,
    PRIMARY KEY ("id")
);
```


## Running the service

> Service will runs on : http://127.0.0.1:9100

> Update main.go with the postgres password chosen in the above step.

> go run main.go


## Return task by UUID

Retreives a task specified by a Uuid.

GET /task/{uuid} 

### Request

Headers : content-type=application/json

### Response

Headers : content-type=application/json

```json
{
    "Id": 19,
    "Uuid": "335979c6-b5b2-4e6d-87da-4bdf8bc5d5f8",
    "Task": "Hello, Task World!",
    "Completed": false,
    "CreatedBy": "AyZee",
    "ModifiedBy": "AyZee",
    "ModifiedOn": "2022-02-22T23:46:19.286798Z"
}
```

## Return all tasks created by user

Retreives all tasks created by the supplied user.

GET /task/{user}

### Request

Headers : content-type=application/json

### Response

Headers : content-type=application/json
 
```json
[
    {
        "Id": 19,
        "Uuid": "335979c6-b5b2-4e6d-87da-4bdf8bc5d5f8",
        "Task": "Hello, Task World!",
        "Completed": false,
        "CreatedBy": "AyZee",
        "ModifiedBy": "AyZee",
        "ModifiedOn": "2022-02-22T23:46:19.286798Z"
    },
    {
        "Id": 20,
        "Uuid": "c55bdb8f-d80b-4179-a9ed-952554a8dce0",
        "Task": "Hello, Task World 2!",
        "Completed": false,
        "CreatedBy": "AyZee",
        "ModifiedBy": "AyZee",
        "ModifiedOn": "2022-02-23T00:09:54.389178Z"
    },
    {
        "Id": 21,
        "Uuid": "68935ecf-d44e-48ab-a1a5-5f616ca2bf4b",
        "Task": "Hello, Task World 3!",
        "Completed": false,
        "CreatedBy": "AyZee",
        "ModifiedBy": "AyZee",
        "ModifiedOn": "2022-02-23T00:09:58.7378Z"
    }
]
```

## Creates task

Create a task to persist to the database. Returns UUID of task.

POST /task/

### Request

Headers : content-type=application/json

```json
{
"Task": "Hello, Task World!",
"Completed": false
}
```

### Response

Headers : content-type=application/json

```
335979c6-b5b2-4e6d-87da-4bdf8bc5d5f8
```

## Update a task

Updates a task by specifying Uuid and completion status. 

POST /task/update

### Request

Headers : content-type=application/json

```json
{
"Uuid": "335979c6-b5b2-4e6d-87da-4bdf8bc5d5f8",
"Completed": true
}
```

### Response

204 - No Content
