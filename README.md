# Bucket List API

Simple item listing API made with Go and PostgreSQL. 
Made to practice the [`go-chi`](https://github.com/go-chi/chi) library and efficient workflow for deployment using Docker and GitHub Actions.

## Dependencies
- Docker
- Docker-compose v1.29+
- Go v1.16+

## ⚙️ How to Run
- Change `.env.sample` to `.env` and update the respective values

- To run from Docker containers
    1. To build and run from containers: `make build`
    2. To bring down containers: `make down`


- If you want to run outside containers:
    1. change `DB_HOST` value to `localhost` in `.env` file
    2. `go mod download`
    3. `go run main.go`



## ☑️ To Do:
- [ ] Tests
- [ ] CI - Github Actions
