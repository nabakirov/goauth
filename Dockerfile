FROM golang:alpine

WORKDIR /go/src/goauth
COPY . .


RUN go get -d -v ./...
RUN go install -v ./...

CMD [ "goauth" ]
EXPOSE 5000