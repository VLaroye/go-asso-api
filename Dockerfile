FROM golang:alpine

RUN apk update && apk add netcat-openbsd

ADD ./wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
EXPOSE 8080

ENTRYPOINT ["/wait-for.sh", "db:5432", "--", "./main"]