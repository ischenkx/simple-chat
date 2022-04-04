# Simple Chat Application

# Architecture
Generally the application is built upon 3 basic structures:

- Repository
- Authorizer
- Event bus

### Repository 
Represents an abstract DataBase. In the current implementation 
the application greedily use repo's resources (a lot of requests).
So, in practice, the repository must take the responsibility of caching.

### Authorizer
Must turn tokens into users' info.

### Event bus
There's a lot of different events happening during the application's lifetime.
Sometimes you want to extend the functionality, so you've got to be able to "intercept" those
events - that's where event bus enters the game.

Note: in practice, an event bus can be used for real time notifications


# Implementation

### Transports
 - HTTP

### Repositories
 - PostgreSQL

### Authorizers
 - JWT

# How to run?
1. Fix the `config.yml` file
2. `go run cmd/web/main.go`

# How to use?
Currently the most User-Friendly way to interact - is a cli tool that's located in `pkg/http/cli`

# Road map
- [ ] Write a simulator for testing
- [ ] Add websockets API (Swirl, Centrifugo)
- [ ] Add caching to postrgres repository (Redis?)
- [ ] Add GraphQL API
- [ ] ...