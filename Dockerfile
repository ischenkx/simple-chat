FROM golang:1.18

WORKDIR /usr/app/src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -v -o /usr/app/bin/app /usr/app/src/cmd/web/main.go

CMD ["/usr/app/bin/app", "-env-config"]