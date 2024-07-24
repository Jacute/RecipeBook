FROM golang:1.22.4


RUN mkdir -p /app

COPY src /app/src

WORKDIR /app/src

RUN go get
RUN go build main.go

CMD ["sh", "-c", "sleep 15 && ./main"]
