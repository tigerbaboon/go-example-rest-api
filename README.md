# go-example-rest-api

- 💻 Example rest api in module structure.

- 💻 Use httprouter with Gin framework.

- 💻 Simple Deploy service on Docker with Dockerfile.

## 🚀 Get Started
```bash
# init project
go mod vendor

# run http
go run .
#or
go run . main.go
```

## 🚀 Route api

- **POST** `/product/create`
  - **Description:** Create a new product.

- **PATCH** `/product/:id`
  - **Description:** Update an existing product by ID.

- **DELETE** `/product/:id`
  - **Description:** Delete a product by ID.

- **GET** `/product/:id`
  - **Description:** Get a product by ID.

- **GET** `/product/list`
  - **Description:** List all products.

## 🚀 Deploy on docker

```bash
# build project to docker images
docker build -t image-name .  

# run container map port with image
docker run --name container-name -p 8080:8080 -d image-name
```

