# **go-jwt-auth**

A simple JWT-based authentication microservice built with Go.

---

## **Main Technology Stack**
- **Programming Language:** Go
- **Database:** SQLite
- **ORM:** GORM
- **Environment Management:** godotenv
- **JWT Library:** [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
- **Encryption:** `crypto/rsa` (RSA256 algorithm for JWT signing)

---

## **Key Technologies Used**
- **Middleware:**
  - CORS (Cross-Origin Resource Sharing)
  - Request logging
  - Authentication
- **JWT Signing Algorithm:** RSA256

---

## **Initial Setup**

1. **Generate RSA Keys**  
   Run the key generation script from the root of the repository:
   ```bash
   go run cmd/rsa-keys-generator/gen_key.go
   ```

2. **Configure Environment Variables**  
   Update the `.env.template` file to specify paths to the generated keys and the SQLite database file. Rename `.env.template` to `.env`.

   Example `.env`:
   ```env
   DB_FILE="test.db"
   PRIVATE_KEY="private_key.pem"
   PUBLIC_KEY="pub_key.pem"
   ```

3. **Run Database Migrations**  
   Apply migrations to create the necessary tables in the database:
   ```bash
   go run cmd/migrations/auto.go
   ```

4. **Start the Application**  
   Launch the microservice:
   ```bash
   go run cmd/main.go
   ```

---

## **Project Structure**

```
├── 📁 cmd
│   ├── 📁 migrations          # Scripts for automatic database migrations
│   ├── 📁 rsa-keys-generator  # RSA key generation script
│   └── 📄 main.go             # Main application file
├── 📁 configs
│   └── 📄 config.go           # Configuration loader
├── 📁 internal                # Internal packages for business logic implementation
│   ├── 📁 auth                # Authentication module (handlers, errors, payload, service)
│   ├── 📁 sqy                 # Test handler with `/sat/hello` endpoint for testing authorization
│   ├── 📁 user                # User package for database operations (create, get, delete)
│   └── 📁 validation          # Custom validators
├── 📁 pkg                     # Reusable packages
│   ├── 📄 db_work             # Database utilities
│   ├── 📄 dependency_injection # Dependency injection utilities
│   ├── 📄 jwt_helper          # JWT utilities
│   ├── 📄 rsa_loader          # RSA key loader
│   ├── 📄 middleware          # Middleware implementations
│   ├── 📄 request             # Request handling utilities
│   └── 📄 response            # Response formatting utilities
├── 📄 LICENSE                 # License file
```

---

## **What’s Interesting About This Project?**

1. **Custom RSA Key Management**  
   Implemented a custom solution for managing RSA keys (`rsa_loader` package).

2. **RSA Keys for JWT**  
   Utilized RSA256 algorithm for signing JWT tokens using private and public keys.

3. **HTTP Response Wrapper**  
   Created a wrapper for `http.ResponseWriter` to capture and log HTTP status codes from requests.

4. **SQLite Instead of PostgreSQL**  
   Chose SQLite for simplicity and lightweight database management during development.

5. **Enhanced Understanding of CORS and DI**  
   Gained practical experience implementing Cross-Origin Resource Sharing (CORS) and Dependency Injection (DI).
