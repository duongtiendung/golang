FROM golang:1.14
WORKDIR /gotest

COPY . .

RUN go get github.com/githubnemo/CompileDaemon


ENTRYPOINT CompileDaemon --build="go build -o main" --command="./main"
