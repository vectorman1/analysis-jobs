FROM golang:1.16

WORKDIR src/jobs
COPY . .

RUN go mod tidy
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["cmd"]