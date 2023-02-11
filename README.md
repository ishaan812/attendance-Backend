# Welcome to Weber!

This is the home of the backend code for our seed application! 

# Requirements
Go v1.18

# Database Setup
The database should be hosted locally on a POSTGRES server, ideally on an empty database, running the server will automatically populate the database.
1.  Install [Postgres](https://www.postgresql.org/download/) locally
2.  Create a local database in your postgres server
3.  Create a .env file in your local root folder and paste the connection string of the created database inside variable 'DATABASE_CONNECTION_STRING'. 
    eg: `DATABASE_CONNECTION_STRING='postgres://<username>:<pwd>@localhost:5432'`
4.  Follow the next steps to run locally. 


# Steps to run locally

1. Run `go run .`
2. Server runs at `http://localhost:9000/`
3. Refer to the [API Documentation](https://docs.google.com/spreadsheets/d/18YYk-Z9BY7OSXgQzrdJ2S9RKgEWv-3_FlDrEU2U6ws0/edit?usp=sharing) to use the API.

# Testing
Import the Postman collection for the API's from the `API_Collection.json` file 

# Dependencies


[github.com/gorilla/handlers](github.com/gorilla/handlers) v1.5.1
[github.com/gorilla/mux](github.com/gorilla/mux) v1.8.0
[github.com/lib/pq](github.com/lib/pq) v1.10.7
[github.com/morkid/paginate](github.com/morkid/paginate) v1.1.4
[gorm.io/driver/postgres](gorm.io/driver/postgres) v1.4.4
[gorm.io/gorm](gorm.io/gorm) v1.24.0

## Important Links
1. How to run a [custom join table](https://github.com/go-gorm/gorm/issues/4051)

