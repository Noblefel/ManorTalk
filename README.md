### ManorTalk
A mini-forum application built using Vue, Go & PostgreSQL

### Dependencies
- [Chi Router](https://github.com/go-chi/chi)
- [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)
- [Migrate](https://github.com/golang-migrate/migrate)
- [GoDotEnv](https://github.com/joho/godotenv)
- [go-validator](https://github.com/go-playground/validator)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [uuid](https://github.com/google/uuid)
- [go-redis](https://github.com/redis/go-redis)

# Installation
```bash
git clone https://github.com/Noblefel/ManorTalk
```  

# Usage (Backend)
### Setup
Navigate inside the directory and download all the dependencies
```bash
cd backend
go mod tidy
go mod download 
``` 

### ENV
Configure the environment variables
```sh
DB_HOST=localhost
DB_NAME=manortalk
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=

REDIS_HOST=localhost
REDIS_PORT=6379 

ACCESS_TOKEN_KEY=access_key
ACCESS_TOKEN_EXP=15m
REFRESH_TOKEN_KEY=refresh_key
REFRESH_TOKEN_EXP=120h
```

### Migrations
Run the migrations if you haven't already, this will create a users table and it's sample users:
```sh
go run cmd/migrate/main.go up
``` 

to revert the migrations:
```sh
go run cmd/migrate/main.go down
``` 

### Start the server
Simply run:
```sh
go run cmd/api/main.go
``` 
(Make sure to have redis server running)

# Usage (Frontend)
Not implemented