FROM golang:1.22.4


WORKDIR /app

COPY config config
COPY database database
COPY handlers handlers
COPY images images
COPY middlewares middlewares
COPY models models
COPY utils utils
COPY go.* .
COPY main.go .

RUN go get
RUN go build main.go

CMD ["sh", "-c", "sleep 15 && ./main"]
