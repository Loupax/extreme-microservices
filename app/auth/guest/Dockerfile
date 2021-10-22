FROM golang AS build
WORKDIR /app
COPY . . 
RUN go mod download
RUN go build /app/app/auth/guest/main.go 

FROM gcr.io/distroless/base
COPY --from=build /app/main /
EXPOSE 8081
ENTRYPOINT ["/main"]