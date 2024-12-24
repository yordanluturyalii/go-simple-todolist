FROM golang:1.24rc1-alpine3.21

WORKDIR /app

COPY . .

ENV DB_USER=root
ENV DB_PASS=root
ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_NAME=todolists

RUN go build -o todolist-app

EXPOSE 8080

RUN chmod +x todolist-app

CMD ["./todolist-app"]