FROM golang:1.19-alpine as builder

WORKDIR /app


# Copy `go.mod` for definitions and `go.sum` to invalidate the next layer
# in case of a change in the dependencies
ADD go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary and prod.json from the builder stage and set it as the default command.
COPY --from=builder /app/main /usr/local/bin/


#COPY --from=builder /app/scripts/ /root/scripts/


CMD ["main"]