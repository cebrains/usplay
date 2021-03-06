############################
# STEP 1 build executable binary
############################
FROM golang:1.14.1-buster as builder
ARG USVC_NAME

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR $GOPATH/src/github.com/FrancescoIlario/usplay

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download \
    && go mod verify

COPY . .

# Build the binary
RUN rm -rf ./bin \
    && make bin PRJ_TARGET=${USVC_NAME} TARGET=cli \
    && mkdir -p /app/bin/ \
    && cp ./bin/${USVC_NAME}/${USVC_NAME}_cli /app/bin/svc_cli

############################
# STEP 2 build an ubuntu image
############################
FROM ubuntu:20.04

# Copy our static executable
COPY --from=builder /app/bin/ /usr/bin/

# Run the produced binary.
ENTRYPOINT ["/usr/bin/svc_cli"]
