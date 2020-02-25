# go-movies
A simple go micro service exposing a movie REST API. 
It is inspired from the online training course [Getting Started with Cloud Native Go](https://www.linkedin.com/learning/getting-started-with-cloud-native-go).

It takes you through the steps of:
- Implementing the micro service from scratch in go.
- Creating a Dockerfile and building a docker image based on that.
- Deploying and running the docker image in a kubernetes cluster.
## Prerequisites
A working installation of go lang.
On mac it can be installed like this:
```bash
brew install go
```
A working installation of [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/).

Clone this repo into your $GOPATH/src folder. 
Use this command to check your current GOPATH:
```bash
go env
```
## Building
After cloning, cd into the go-movies folder and execute:
```bash
go build
```
## Testing
The app api can be tested like this:
```bash
go test ./api -v
```

## Running
Once built and tested, execute the app like this:
```bash
./go-movies
```
Now you can use for instance Postman to send requests to the running app.

## Build Docker Image
Start your minikube unless it is already started. 
Also, make sure to setup the docker environment:
```bash
eval $(minikube docker-env)
```
Build a docker image of the app like this:
```bash
docker build -t cloud-native-go-movies:1.0.0 .
```
## Running Docker Image
There are many ways to run the built image. 
To run it as a daemon called "movies", use this command:
```bash
docker run --name movies -d -p 8080:8080 cloud-native-go-movies:1.0.0
```
To stop and remove the daemon:
```bash
docker stop movies
docker rm movies
```
While it is running you can test sending Postman requests to it. 
Note that the app will be running on the ip specified by the $DOCKER_HOST environment variable rather than localhost.

To test that the app is able to serve requests on a different port, you can add the MOVIE_PORT environment variable using the -e flag when starting the daemon:
```bash
docker run --name movies -d -e "MOVIE_PORT=9090" -p 9090:9090 cloud-native-go-movies:1.0.0
```

