# What is this 

An API to hold where users have been in the world.

# Routes 

POST /v1/users/create               -> create a user 
GET /v1/users/{user_id}             -> get specific user
GET /v1/countries                   -> get all countries
GET /v1/regions                     -> get all regions
GET /v1/user/countries             -> get countries user has visited or not 
GET /v1/user/suggestions           -> get country suggestions that a user might like
POST /v1/users/visit                -> record that a user has visited a country
GET /v1/refresh-maps                -> refresh maps table using open source API

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
