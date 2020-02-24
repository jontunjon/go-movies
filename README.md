# go-movies

## Prerequisites
A working installation of go lang.
On mac it can be installed like this:
```bash
brew install go
```
A working installation of [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/).

Clone this repo into your $GO_PATH/src folder. 
Use this command to check your current GO_PATH:
```bash
go env
```
## Building
After cloning, cd into the go-movies folder and execute:
```bash
go build
```
## Testing
TBD

## Running
Once built and tested, execute the app like this:
```bash
./go-movies
```
Now you can use for instance Postman to send requests to the running app.

## Build Docker Image
Before building the image, make sure to setup the docker environment:
```bash
eval $(minikube docker-env)
```
Build a docker image of the app like this:
```bash
docker build -t cloud-native-go-movies:1.0.0 .
```

