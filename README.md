# OneCV_Internship_test
Source code for OneCV Technical Test

## How To Run
## Prerequisites
```
Have go installed on your machine
Ensure that you have a local database instance of mysql for the connections
```
### Install neccesary modules
```go mod download```
### Set up an .env file
```
Create a .env file under the root folder of the directory and paste/edit the configuration below to your needs
```
```
PORT=3000 # Your port which you would like to run the api on
DB_URL="user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
# 'user' is the username of your database
# 'pass' is the password of your database
# '127.0.0.1' URL of your database
# ':3306' your database port
# 'dbname' your database instance name
```
### Migrate your database with tables and sample data
```
go run migrate/migrate.go
```

## Run API
```
go run main.go

OR

CompileDaemon -command="./OneCV_Internship_test" # Run the server locally with updates
```

## Running Test cases
```
cd controllers
go test
```

