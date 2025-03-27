# go-jwt-auth
A simple jwt auth microservice on go.

# Stack
Go, sqlite, gorm, godotenv

# Initial setup
Create .env file and determine db file

Example:
```.env
DbFile="test.db"
```

Run migrations:

```
go run cmd/migrations/auto.go
```

