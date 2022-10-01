
# Decription this project
 This is example REST API using GoLang, not use framework and database.

# How to run this project

1. Create an .env file whose contents are based on the .env.example file
2. run the project

```
go run main.go
```
# Testing api with postman

```
import the file in Postman located at ./Students.postman_collection.json
```


# REST API Specification

### GROUP: Student

- [1] - Create
- [POST] : {root.api}/api/v1/student

```json

Request Body:
{
    "name": "Budi",
    "age": 5
}

Response:
{
    "meta": {
        "message": "success create student",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 1,
        "name": "Budi",
        "age": 5
    }
}
```

- [2] - List
- [GET] : {root.api}/api/v1/students
- [GET] : {root.api}/api/v1/students?page=1&size=10

```json

Request Param:
-default page = 1
-default size = 10

Response:
{
    "meta": {
        "message": "success get list students",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "id": 1,
            "name": "Budi",
            "age": 5
        },
        {
            "id": 2,
            "name": "Budi 2",
            "age": 5
        },
        {
            "id": 3,
            "name": "Budi 3",
            "age": 5
        },
        {
            "id": 4,
            "name": "Budi 4",
            "age": 5
        },
        {
            "id": 5,
            "name": "Budi 5",
            "age": 9
        }
    ],
    "page": {
        "size": 10,
        "total_data": 5,
        "total_page": 1,
        "current": 1
    }
}

```

- [3] - Get By ID
- [GET] : {root.api}/api/v1/student/{id}

```json

Response:
{
    "meta": {
        "message": "success get student",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 2,
        "name": "Budi 2",
        "age": 5
    }
}

```

- [4] - Delete By ID
- [DELETE] : {root.api}/api/v1/student/{id}

```json

Response:
{
    "meta": {
        "message": "success delete student",
        "code": 200,
        "status": "success"
    },
    "data": null
}

```

- [5] - Update By ID
- [PUT] : {root.api}/api/v1/student/{id}

```json

Request Body:
{
    "name": "Budi Kurniawan",
    "age": 15
}

Response:
{
    "meta": {
        "message": "success get student",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 2,
        "name": "Budi Kurniawan",
        "age": 15
    }
}
