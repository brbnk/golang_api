FROM golang:1.13.4
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
EXPOSE 7777