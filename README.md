### ManorTalk
A mini-forum application built using Vue, Go & PostgreSQL

### Dependencies
- [Chi Router](https://github.com/go-chi/chi)
- [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)
- [Migrate](https://github.com/golang-migrate/migrate)
- [GoDotEnv](https://github.com/joho/godotenv)

<br>

# Installation
```bash
git clone https://github.com/Noblefel/ManorTalk
``` 
<br>

# Usage (Backend)
### Setup
Navigate inside the directory and download all the dependencies
```bash
cd backend
go mod tidy
go mod download 
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

# Usage (Frontend)
Not implemented