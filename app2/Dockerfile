FROM golang:1.14
WORKDIR /app2

COPY . .

RUN go get github.com/githubnemo/CompileDaemon


ENTRYPOINT CompileDaemon --build="go build -o main2" --command="./main2"