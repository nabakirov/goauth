FROM golang:alpine

WORKDIR /go/src/goauth
COPY . .

RUN apk --update add --no-cache git

RUN go get github.com/gorilla/mux && go get github.com/gorilla/handlers

RUN go build goauth

ENTRYPOINT ["./goauth"]
EXPOSE 5000