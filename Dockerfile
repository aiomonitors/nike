FROM golang:latest

WORKDIR /usr/app/backend

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

RUN ls

CMD ["./main"]