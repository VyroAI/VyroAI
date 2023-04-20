# Build stage
FROM golang:latest AS Builder

ARG github_token

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./config.prod.yml .

COPY . /app

RUN CGO_ENABLED=0 go build main.go

FROM alpine:latest as Server

WORKDIR /app

COPY --from=Builder /app/main ./
COPY --from=Builder /app/config.prod.yml ./

RUN chmod +x ./main

CMD [ "./main" ]
