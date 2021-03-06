FROM golang:1.13.8
MAINTAINER jontunjon
RUN eval $(go env)
ENV SOURCES=$GOPATH/src/go-movies
COPY . $SOURCES
RUN cd $SOURCES && CGO_ENABLED=0 go install
ENV MOVIE_PORT=8080
EXPOSE 8080
ENTRYPOINT go-movies