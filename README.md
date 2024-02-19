### ManorTalk
A mini-forum application built in Go 1.20

### Dependencies
- [Chi Router](https://github.com/go-chi/chi)
- [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)
- [Migrate](https://github.com/golang-migrate/migrate)
- [GoDotEnv](https://github.com/joho/godotenv)
- [go-validator](https://github.com/go-playground/validator)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [uuid](https://github.com/google/uuid)
- [go-redis](https://github.com/redis/go-redis)
- [slug](https://github.com/gosimple/slug)

### TODO
1. Create, edit posts (frontend)
2. Refactor JSON response to separate pagination metadata from the main data.
3. Refactor unit test.
4. Improve database queries in post repo.
5. Fix migration to set from specific version/steps, instead of migrating up/down everything.

## Installation
```bash
git clone https://github.com/Noblefel/ManorTalk
```  

# Usage (with Docker)

Build our image and start the containers:

```sh
docker compose up --build 
```

### Required Environment Variables

| Key | Sample |
| -------- | ------- |
| API_PORT | 8080 |
| DB_HOST | localhost |
| DB_NAME | manortalk |
| DB_PORT | 5432 |
| DB_USER | postgres |
| DB_PASSWORD |  |
| REDIS_HOST | localhost |
| REDIS_PORT | 6379 |
| ACCESS_TOKEN_KEY | access_key |
| REFRESH_TOKEN_KEY | refresh_key |

# Usage (Local)
### 1. Backend
### Setup
Navigate inside the directory and download all the dependencies
```bash
cd backend
go mod tidy
go mod download 
``` 

### .env
Configure the environment variables inside the backend directory 

### Migrations
Run the migrations if you haven't already, this will create a users table and it's sample users:
```sh
go run cmd/migrate/main.go -production=false up
``` 

to revert the migrations:
```sh
go run cmd/migrate/main.go -production=false down
``` 

### Start the server
Simply run:
```sh
go run cmd/api/main.go -production=false
``` 
(Make sure to have redis server running)

### 2. Frontend
Navigate inside the directory and download all the dependencies
```bash
cd frontend
npm install 
```

Start development
```sh
npm run dev 
```