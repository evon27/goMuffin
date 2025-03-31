FROM golang:1.24.1

RUN mkdir /app
WORKDIR /app

COPY ./commands .
COPY ./configs .
COPY ./databases .
COPY ./handler .
COPY ./utils .
COPY ./components .
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build -o build/goMuffin git.wh64.net/muffin/goMuffin

ENTRYPOINT [ "./build/goMuffin" ]