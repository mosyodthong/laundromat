FROM golang:1.20.3-alpine3.16

WORKDIR /app

COPY . .
RUN go mod download -x
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go mod tidy
ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"