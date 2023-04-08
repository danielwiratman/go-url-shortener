FROM golang:1.20.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /godocker
EXPOSE 9090
CMD [ “/godocker” ]