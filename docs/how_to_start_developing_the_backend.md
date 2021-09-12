# How to start developing the backend

## Parameters needed for .env

```
PORT
MONGO_URI
MONGO_DB
```

PORT = port to run the backend server
MONGO_URI = the URI of MongoDB Atlas
MONGO_DB = the database within MongoDB Atlas to connect to

## How to start the server

In the root of the folder, run `go run main.go` to start the development server

## How to start the server with docker

1. In the root of the project, run `docker build -t ice-cream-backend .` to build the docker image

2. After the build is successful, run `docker run -p 3001:3001 -it ice-cream-backend` to start the application on `localhost:3001`
