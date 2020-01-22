FROM golang:1.13.6 AS build-env
RUN mkdir /work
WORKDIR /work
COPY . .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -o app

FROM alpine
RUN mkdir /work
WORKDIR /work
COPY --from=build-env /work/app /work/app

ENTRYPOINT ["/work/app"]
