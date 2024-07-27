FROM golang:1.23rc2


RUN mkdir -p /app

COPY src /app/src

WORKDIR /app/src

RUN go get
RUN go build main.go

CMD ["sh", "-c", "sleep 15 && ./main"]
