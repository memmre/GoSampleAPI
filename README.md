# GoSampleAPI

GoSampleAPI is a RESTful API built with Golang, designed for handling user authentication and event management. This project follows a modular structure to ensure scalability and maintainability.

## Features
- User authentication (Sign up, Sign in)
- JWT-based authentication
- Secure password hashing
- CRUD operations for events
- Middleware for authorization

## Project Structure
```
GoSampleAPI/
├── database/
│   └── database.go
├── middlewares/
│   └── auth.go
├── models/
│   ├── event.go
│   └── user.go
├── routes/
│   ├── events.go
│   ├── registrations.go
│   ├── routes.go
│   └── users.go
├── test/
│   ├── create-event.http
│   ├── create-registration.http
│   ├── delete-event.http
│   ├── delete-registration.http
│   ├── get-event.http
│   ├── get-events.http
│   ├── sign-in.http
│   ├── sign-up.http
│   └── update-event.http
├── utilities/
│   ├── hash.go
│   └── jwt.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Installation

### Prerequisites
- [Go](https://go.dev/doc/install) installation
- SQLite3 or any other preferred database

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/GoSampleAPI.git
   ```
2. Navigate into the project directory:
   ```sh
   cd GoSampleAPI
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
5. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

### Authentication
- **POST** `/signIn` - Login and get a JWT token
- **POST** `/signUp` - CreateRegistration a new user

### Events
- **POST** `/events` - Create a new event (Authenticated)
- **GET** `/events` - Retrieve all events
- **GET** `/events/{id}` - Retrieve a specific event
- **PUT** `/events/{id}` - Update an event (Authenticated)
- **DELETE** `/events/{id}` - Delete an event (Authenticated)
- **POST** `/events/{id}/register` - Create an event registration (Authenticated)
- **DELETE** `/events/{id}/register` - Delete an event registration (Authenticated)

## Testing API Requests
The `test` directory contains HTTP request samples that can be tested using an API client like [REST Client for VS Code](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

## Security
- Passwords are securely hashed using bcrypt (`utilities/hash.go`).
- JWT is used for authentication (`utilities/jwt.go`).
- Authorization middleware ensures restricted access to protected routes (`middlewares/auth.go`).

## Contributing
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Create a Pull Request
