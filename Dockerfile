FROM golang:1.24.1
WORKDIR /usr/src/app/

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy
# RUN go mod init github.com/feysalhassan1/Go-API
# RUN go get github.com/gofiber/fiber/v2 

# Create tmp directory and set permissions
RUN mkdir -p tmp && chmod 777 tmp

CMD ["air", "-c", "air.toml"]