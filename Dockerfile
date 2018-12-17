FROM golang:alpine as builder
RUN mkdir -p $GOPATH/src/github.com/vitr/docsify-server
WORKDIR $GOPATH/src/github.com/vitr/docsify-server
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app
FROM scratch
COPY --from=builder /app /app
EXPOSE 8000
ENTRYPOINT ["/app"]