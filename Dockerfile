FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get github.com/lib/pq

RUN go get github.com/joho/godotenv


RUN go build -o main .

EXPOSE 80

CMD [ "./main" ]