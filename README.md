# Have fun with GIN web framework - Recipe Management API

This is a simple recipe management API built with Go and the Gin framework. It allows you to create, read, update, and delete recipes stored in an in-memory storage.

## Features

- Add a new recipe
- Get a recipe by its slug
- List all recipes
- Update a recipe
- Delete a recipe by its slug

## Setup

### Prerequisites

- [Go](https://golang.org/dl/) installed
- [Gin](https://github.com/gin-gonic/gin) web framework
- [gosimple/slug](https://github.com/gosimple/slug) library

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/recipe-management-api.git
    cd recipe-management-api
    ```

2. Install the required dependencies:
    ```sh
    go get -u github.com/gin-gonic/gin
    go get -u github.com/gosimple/slug
    ```

3. Build and run the application:
    ```sh
    go run main.go
    ```

The server will start at `localhost:8080`.

## Endpoints

### Add a New Recipe

- **URL**: `/recipes`
- **Method**: `POST`
- **Headers**: `Content-Type: application/json`
- **Body**:
    ```json
    {
      "name": "user friendly name of my favorite cookies",
      "title": "The Modern Sound of Betty Carter",
      "author": "Betty Carte"
    }
    ```

### List All Recipes

- **URL**: `/recipes`
- **Method**: `GET`

### Get a Recipe by ID

- **URL**: `/recipes/:id`
- **Method**: `GET`

### Update a Recipe

- **URL**: `/recipes/:id`
- **Method**: `PUT`
- **Headers**: `Content-Type: application/json`
- **Body**:
    ```json
    {
      "name": "Updated Name",
      "title": "Updated Title",
      "author": "Updated Author"
    }
    ```

### Delete a Recipe by ID

- **URL**: `/recipes/:id`
- **Method**: `DELETE`

## Example Requests


```http
POST localhost:8080/recipes
Content-Type: application/json

{
  "name": "1",
  "title": "The Modern Sound of Betty Carter",
  "author": "Betty Carte"
}
```

```http
DELETE localhost:8080/recipes/1
```

```http
GET localhost:8080/recipes
```

```http
GET localhost:8080/recipes/{id}
```