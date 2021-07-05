FROM golang:1.16-alpine AS build

WORKDIR /build
COPY ./ ./
ENV CGO_ENABLED 0
RUN go mod download && go mod verify && go build .

FROM scratch
WORKDIR /embassy
COPY --from=build /build/embassy /embassy/embassy

ENTRYPOINT [ "/embassy/embassy" ]
