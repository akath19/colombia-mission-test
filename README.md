# Colombia Mission Technical Test
This repo holds all files requested by the Colombia Mission Technical Operations Test

## Configuration
The following environment variables must be passed to the Docker container for the app to start properly

| Property | Description |
| -------- | ----------- |
| POSTGRES_ADDR | PostgreSQL DB to connect to, this address must not contain a protocol |
| POSTGRES_PORT | PostgreSQL DB port to connect to |
| POSTGRES_USER | PostgreSQL user to connect as |
| POSTGRES_PASS | PostgreSQL password for the above user |
| POSTGRES_DB | PostgreSQL database to use |
| HTTP_PORT | http port to expose service in |

Please remember to create a user with `SELECT, INSERT, UPDATE, DELETE & CONNECT` privileges to the database selected in `POSTGRES_DB`

## How To Run
1. Clone this repo
1. Download & install golang from [here](https://golang.org/)
1. Run `go get` to download all dependencies
1. Run `go build` to create the final executable
1. Run `colombia-mission-test` executable after setting up the above environment variables & PostgreSQL instance