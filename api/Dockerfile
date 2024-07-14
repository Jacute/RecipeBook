FROM golang:1.22.4


WORKDIR /app

COPY . .
RUN go get
RUN go build main.go

CMD ["sh", "-c", "sleep 10 && ./main"]
