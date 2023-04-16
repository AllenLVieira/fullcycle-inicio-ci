FROM golang:1.19

WORKDIR /app

RUN go mod init AllenLVieira/fullcycle-inicio-ci

COPY . .

RUN go build -o math

CMD ["./math"]