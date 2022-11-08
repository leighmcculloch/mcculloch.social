FROM golang AS go
WORKDIR /go/src/
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go install

FROM scratch
COPY --from=go /go/bin/mcculloch.social /mcculloch.social
ENTRYPOINT ["/mcculloch.social"]
