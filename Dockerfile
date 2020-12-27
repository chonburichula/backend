FROM golang as backend
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY main.go ./
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
