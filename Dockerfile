FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/lorenzomrt/content-insight
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/lorenzomrt-cr-api ./cmd/api

FROM scratch
COPY --from=build /go/bin/lorenzomrt-cr-api /go/bin/lorenzomrt-cr-api
ENTRYPOINT [ "/go/bin/lorenzomrt-cr-api" ]