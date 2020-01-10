FROM go-hello-world-build:latest as build
FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=build /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 8080