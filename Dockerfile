FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -tags netgo -ldflags '-s -w' -o whypfs-api

EXPOSE 1313

CMD [ "/whypfs-gateway" ]