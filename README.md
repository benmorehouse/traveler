# What is this 

An API to hold where users have been in the world and suggest to them where to go next based on where they have been.

# Routes 

GET    /v1/status                --> status of the api

POST   /v1/users/create          --> create a user

GET    /v1/user                  --> get the user

GET    /v1/countries             --> countries

GET    /v1/user/suggestions      --> get suggestions on where the user should go next

POST   /v1/user/visit            --> user has performed a visit

POST   /v1/internal/refresh      --> internal tool to refresh countries in database

# Structure

traveler/
│
├── cmd/                      # Application entry points
│   └── myapp/                # Main application
│       └── main.go           # Main program file
│
├── pkg/                      # Library code used by your app
│   ├── api/                  # API domain logic
│   │   ├── handler/          # HTTP handlers (controllers)
│   │   ├── middleware/       # HTTP middlewares
│   │   └── response/         # API response structures and utilities
│   │
│   ├── model/                # Data models and business logic
│   ├── repo/                 # Database interactions (repository layer)
│   ├── service/              # Business logic (service layer)
│   └── util/                 # Utility functions and common helpers
│
├── config/                   # Configuration files and parsers
│   └── config.go             # Configuration struct and loader
│
├── migrations/               # Database migration files
│
├── Dockerfile                # Dockerfile for building the application
├── Makefile                  # Makefile for automating common tasks
├── go.mod                    # Go module file
└── go.sum                    # Go module checksum file
