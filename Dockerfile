FROM golang:latest

WORKDIR /usr/src/app

COPY . .

WORKDIR /usr/src/app/Pruthvi

CMD ["go", "run", "main.go"]
