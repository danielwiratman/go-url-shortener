FROM golang:1.20.3
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /godocker
EXPOSE 9090
CMD [ “/godocker” ]