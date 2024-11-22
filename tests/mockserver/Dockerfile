FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download && go mod verify
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod go build -v -o /usr/local/bin/mockserver .
ENTRYPOINT ["mockserver"]
