FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/api/cmd

RUN go build -o api .

CMD [ "./api" ]