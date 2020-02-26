# go-movies
A simple go micro service exposing a movie REST API. 
It is inspired from the online training course [Getting Started with Cloud Native Go](https://www.linkedin.com/learning/getting-started-with-cloud-native-go).

It takes you through the steps of:
- Implementing the micro service from scratch in go.
- Creating docker images in a number of different ways.
- Installing the docker image in a kubernetes cluster.
## Prerequisites
A working installation of go lang.
On mac it can be installed like this:
```bash
brew install go
```
A go-capable IDE of your choice such as GoLand or Visual Studio Code.

A working installation of [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/).

Clone this repo into your $GOPATH/src folder. 
Use this command to check your current GOPATH:
```bash
go env
```
## Building Go app
After cloning, cd into the go-movies folder and execute:
```bash
go build
```
## Testing Go app
The app api can be tested like this:
```bash
go test ./api -v
```

## Running Go app
Once built and tested, execute the app like this:
```bash
./go-movies
```
Now you can use for instance Postman to send requests to the running app.

## Building Docker Image
Start your minikube unless it is already started. 
Also, make sure to setup the docker environment:
```bash
eval $(minikube docker-env)
```
Build a docker image of the app like this:
```bash
docker build -t cloud-native-go-movies:1.0.0 .
```
This will build a docker image based on the details in [Dockerfile](./Dockerfile).

To check the result:
```bash
docker images

REPOSITORY                                TAG                 IMAGE ID            CREATED             SIZE
cloud-native-go-movies                    1.0.0               c1e044b38ae1        1 minute ago        827MB
```
## Running Docker Image
There are many ways to run the built image. 
To run it as a daemon in a container called "movies", use this command:
```bash
docker run --name movies -d -p 8080:8080 cloud-native-go-movies:1.0.0
```
To stop and remove the container:
```bash
docker stop movies
docker rm movies
```
While it is running you can test sending Postman requests to it. 
Note that in this case the app will be running on the ip specified by the $DOCKER_HOST environment variable rather than localhost.

To test that the app is able to serve requests on a different port, you can add the MOVIE_PORT environment variable using the -e flag when starting the daemon:
```bash
docker run --name movies -d -e "MOVIE_PORT=9090" -p 9090:9090 cloud-native-go-movies:1.0.0
```
## Docker Compose
Docker compose provides another way of building and running docker images.
In this example a more minimalistic docker image will be built and an nginx service will also be included in the image.
Check out the [Dockerfile-minimal](./Dockerfile-minimal) and [docker-compose.yaml](./docker-compose.yaml) files for details.

In order to build this minimal image, we must first build the go app for the target architecture. 
In this case linux amd64:
```bash
GOOS=linux GOARCH=amd64 go build
```
Once that is done, a new docker image can be built simply by running this command:
```bash
docker-compose build
```
This image is considerably smaller compared to the one previously built:
```bash
docker images

REPOSITORY                                TAG                 IMAGE ID            CREATED              SIZE
cloud-native-go-movies                    1.0.0               09189107cb1a        2 minutes ago        827MB
cloud-native-go-movies                    1.0.1               5048b98580b4        About a minute ago   21MB
```
Again we can run it as a daemon but now via a docker-compose command:
```bash
docker-compose up -d
```
You can test that the app serves requests just as before using Postman and also that the nginx service is available by going to http://<DOCKER_HOST>:8080 in your browser.

To stop and remove the container:
```bash
docker-compose down
```

You can test that the app serves requests just as before using Postman and also that the nginx service is available by going to http://<DOCKER_HOST>:8080 in your browser.

## Pushing Docker Image to Remote Repository
We can push the docker image to a remote repository using these steps:
1. Create a suitable repository in the docker registry of your choice.
   In this example the repository **jonas.thungren/cloud-native-go-movies** in registry **dtr.digitalroute.com** is used.
2. Login to the repository in the shell:
```bash
docker login dtr.digitalroute.com
```
3. Tag the docker image for the remote respository:
```bash
docker tag cloud-native-go-movies:1.0.1 dtr.digitalroute.com/jonas.thungren/cloud-native-go-movies:1.0.1
```
4. Push the docker image to the remote repository:
```bash
docker push dtr.digitalroute.com/jonas.thungren/cloud-native-go-movies:1.0.1
```

## Installing in Kubernetes Using Helm
A [helm configuration](./helm) for the docker image has been created to facilitate installation in kubernetes (in this case minikube).
Installation based on this config is simply done using this command:
```bash
cd helm
helm install movies .
```
Once the installation is done, check that the expected pod has been created and that it is running:
```bash
kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
cloud-native-go-movies   1/1     Running   0          8s
```
Once again we can check that the app is serving requests using Postman. Just make sure to start port forwarding first in order to access the pod port from the local machine:
```bash
kubectl port-forward cloud-native-go-movies 8080:8080
```
To uninstall, simply use this command:
```bash
helm uninstall movies
```
This will stop and delete everything that was previously installed using the helm configuration.

We can also verify that the port that the app is serving requests on is configurable by updating the **MoviePort** value in the [values.yaml](./helm/values.yaml) and then running the install command again.
This time the portforwarding should be made against whatever the MoviePort value was set to.
