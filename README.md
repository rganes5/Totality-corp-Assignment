# Totality-corp-Assignment

#### This has a api-gateway and a user service. Since it is a simple project, a mono-repo method rather than polyrepo is used. 
#### This project comprises a set of microservices built using the Go programming language. 
#### It adheres to Clean Architecture principles, utilizes GRPC for communication, and uses Swagger for API documentation.
#### The API is split into two main microservices: an API Gateway and one user microservice.

### Additional Features and Libraries


The project utilizes the following packages:
1. [GIN](github.com/gin-gonic/gin): A web framework written in Go that combines high performance with an API similar to Martini.
2. [Wire](https://github.com/google/wire): A code generation tool for dependency injection, making it easier to connect components.
3. [Viper](https://github.com/spf13/viper): A configuration solution for Go applications, supporting various formats and 12-Factor app principles.
4. [swag](https://github.com/swaggo/swag) with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files): Converts Go annotations to Swagger Documentation 2.0 for API documentation.
5. [Clean Code Architecture](https://www.freecodecamp.org/news/a-quick-introduction-to-clean-architecture-990c014448d2/): Implemented to achieve separation of concerns and maintainability.
6. [gRPC](https://grpc.io/): Employed for inter-microservice communication due to its lightweight and efficient protocol.
7. [mockgen](https://github.com/golang/mock): The project utilizes mockgen, a code generation tool that automatically creates mocks for the defined interfaces. This is particularly useful for isolating units under test and validating their behavior without the need for real implementations. 
8. Loose Coupling: Designed for independent development, testing, and deployment of microservices.


# To Run Project

## Run using Docker-Compose

##### Setup env file for docker-compose (look up the .env.example for your reference)
create .env file on the project root dir and add the below envs on it
```.env
#example

PORT=:3000
USER_SVC_PORT=:50051
USER_SVC_URL=user-service:50051

```
#### Run Docker Containers
```
docker compose up
```
#### To See The API Documentation 
http://localhost:3000/swagger/index.html


# TO RUN MANUALLY

#### Clone The Repository
```
git clone https://github.com/rganes5/Totality-corp-Assignment.git
```
#### Checkout To Project Directory
```
cd ./Totality-corp-Assignment
```
### . Install Dependencies
Install the required dependencies using either the provided Makefile command or Go's built-in module management:
```bash
# Using Corresponding Makefile in each Microservices
# OR using Go
go mod tidy
```

```
cd ./api-gateway
```
go mod tidy
make proto
make wire
make swag

```bash
make run
```

```
cd ./user-service
```
go mod tidy
make proto
make wire
make swag
```bash
make run
```
## To Test The Application
### Generate Mock Files
```bash
go install github.com/golang/mock/mockgen@v1.6.0
```
```bash
go get github.com/golang/mock/gomock
```

```bash
go test ./... -cover
```

#### To See The API Documentation 
http://localhost:3000/swagger/index.html



