# FeatureToggle Golang Project

A feature toggle is a technique in software development that is used to hide, enable or disable the feature during runtime.

This is a light, simple and fast microservice application builded with golang. It uses a postgres database to store the features. And have users and authentication control.

### How to use

To run the application, you just need to use the docker-compose to build the infrastructure and then run the application on port 8080 using the command below.

go run ./cmd/main.go

### How it works