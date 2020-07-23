FROM golang:1.14
WORKDIR /web

COPY . .

RUN go get github.com/githubnemo/CompileDaemon


ENTRYPOINT CompileDaemon --build="go build -o main" --command="./main"