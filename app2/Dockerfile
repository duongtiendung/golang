FROM golang:1.14

RUN mkdir /app2
# We copy everything in the root directory
# into our /app directory
ADD . /app2
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app2
# we run go build to compile the binary
# executable of our Go program
RUN go build -o main2 .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app2/main2"]