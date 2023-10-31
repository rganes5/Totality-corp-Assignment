# Totality-corp-Assignment
This has a api-gateway and a user service. Since it is a simple project, a mono-repo method rather than polyrepo is used. 

# To Run Project

#### Clone The Repository
```
git clone https://github.com/rganes5/Totality-corp-Assignment.git
```
#### Checkout To Project Directory
```
cd ./Totality-corp-Assignment
```
## Run using docker-compose

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
