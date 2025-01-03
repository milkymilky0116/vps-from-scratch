FROM golang:latest AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o simple_api .

FROM build

WORKDIR /app

COPY --from=build /app/simple_api .

RUN chmod +x ./simple_api

EXPOSE 8080 
CMD ["./simple_api"]
