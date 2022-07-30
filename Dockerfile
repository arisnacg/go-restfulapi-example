FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-restfulapi-example

EXPOSE 3000

CMD [ "/go-restfulapi-example"]
