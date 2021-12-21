FROM golang:1.17
LABEL maintainer="Streamer272 <admin@streamer272.com>"

WORKDIR /home/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["ddos"]

CMD ["--help"]
