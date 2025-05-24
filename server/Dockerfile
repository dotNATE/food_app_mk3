FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o food_app

EXPOSE 8080

ENTRYPOINT [ "./food_app" ]