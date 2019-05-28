############################
# STEP 1 build executable binary
############################

FROM golang:1.12-alpine as builder

# Install SSL ca certificates.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache ca-certificates


COPY . $GOPATH/src/github.com/ElegantCreationism/go-hoover/
WORKDIR $GOPATH/src/github.com/ElegantCreationism/go-hoover/


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/svc -ldflags="-s -w"

# Using go mod.

# Build the binary


############################
# STEP 2 build a small image
############################

FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /go/bin/svc /svc

# Port on which the service will be exposed.
EXPOSE 8080

# Run the svc binary.
CMD ["./svc"]

