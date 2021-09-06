FROM golang

WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 go build -o /bin/app cmd/main.go

FROM alpine

WORKDIR /bin
COPY --from=0 /bin/app .
COPY screener_files /bin/screener_files

ENTRYPOINT ["app"]
