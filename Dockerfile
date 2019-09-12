ARG GO_VERSION=1.12

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache ca-certificates git tzdata

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
# copying images files to the root
COPY ./images /images 

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .


# FINAL CONTAINER
FROM scratch AS final

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app /app

# so our images are available in the final container
COPY --from=builder /images /images

EXPOSE 80

ENTRYPOINT ["/app"]