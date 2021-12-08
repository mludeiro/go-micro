FROM golang:1.17-alpine3.15

# Set destination for COPY
WORKDIR /app
COPY . ./

# Download Go modules and build
RUN apk --no-cache add build-base \
    && go mod download -x \
    && go build -o /go-micro -x\
    && rm -rf /go \
    && apk del build-base 

EXPOSE 9000

# Run
CMD [ "/go-micro" ]
