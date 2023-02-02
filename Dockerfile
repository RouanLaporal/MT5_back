FROM golang:1.18 as BUILDER

# Active le comportement de module indépendant
ENV GO111MODULE=on

# Je vais faire une build en 2 étapes
# https://dave.cheney.net/2016/01/18/cgo-is-not-go
ENV CGO_ENABLED=0
ENV GOOS=$GOOS
ENV GOARCH=$GOARCH

WORKDIR /go_app
COPY ./go_app .
RUN go mod download \
    && go mod verify \
    && go build -o /build/buildedApp main/main.go

FROM scratch as FINAL

WORKDIR /main
COPY --from=BUILDER /build/buildedApp .

ENTRYPOINT ["./buildedApp"]