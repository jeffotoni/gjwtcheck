FROM golang:1.17 as builder
WORKDIR /go/src/gjwtcheck
COPY . . 

ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build --trimpath -ldflags="-s -w" -o gjwtcheck main.go
RUN cp gjwtcheck /go/bin/gjwtcheck

FROM alpine:latest as builder2
RUN apk add --no-cache upx
RUN apk add --no-cache tzdata
ENV TZ America/Sao_Paulo

# Relocate the timezone file
RUN mkdir -p /config/etc && mv /etc/timezone /config/etc/ && ln -s /config/etc/timezone /etc/

COPY --from=builder /go/bin/gjwtcheck /go/bin/gjwtcheck
WORKDIR /go/bin
RUN upx gjwtcheck
RUN apk del --no-cache upx

FROM scratch
# Copy our static executable.
COPY --from=builder2 /go/bin/gjwtcheck /
# Run the hello binary.
EXPOSE 8080
ENTRYPOINT ["/gjwtcheck"]
