# Go REST API Example

A simple REST API build by GO langunge for research and learning purpose

## Installation and Run

Download this project from Github

```bash
go get github.com/arisnacg/go-restfulapi-example
```

Make `.env` like down below

```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=yourdbusername
DB_PASSWORD=yourdbpassword
DB_NAME=go_example
API_SECRET=yoursecretstring
TOKEN_HOUR_LIFESPAN=1
```

Build and run the project

```bash
cd go-restfulapi-example
go build
./go-restfulapi-example

```

API Endpoint: `http://localhost:3000`

## Structure

```
├── main.go
├── controllers
│   └── auth.go // controller for authentication endpoints
├── middlewares
│   └── middlewares.go // auth middlewares
├── models
│   ├── setup.go // setup for database
│   └── user.go // user model
└── utils
    └── token
        └── token.go // JWT utils
```

## API

### Public

- `POST` `/api/register` : Register a user
- `POST` `/api/login` : Login with username and password to get JWT token

### Authenticated

- `GET` `/api/admin/me` : Get current login user
