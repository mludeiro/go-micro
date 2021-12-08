FROM golang:1.16-alpine

RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 

# Set destination for COPY
WORKDIR /app
COPY . ./

# Download Go modules and build
RUN go mod download &&  go build -o /go-micro && rm -rf /go

EXPOSE 9000

# Run
CMD [ "/go-micro" ]
