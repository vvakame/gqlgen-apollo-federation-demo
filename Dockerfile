FROM golang:1.17

ARG SERVICE_NAME

WORKDIR /project

COPY . ./

RUN go build -o server ./${SERVICE_NAME}/server/server.go

CMD ["/project/server"]
