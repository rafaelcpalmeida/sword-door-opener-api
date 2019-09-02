FROM golang:1.12

WORKDIR /app
EXPOSE 40000 8080

RUN go get github.com/derekparker/delve/cmd/dlv

ADD main.go .

CMD [ "dlv", "debug", "--listen=:40000", "--headless=true", "--api-version=2", "--log" ]
