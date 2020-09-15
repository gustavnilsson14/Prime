FROM golang:1.11
RUN go get -u -v github.com/gin-gonic/gin
RUN go get -u -v github.com/gin-gonic/contrib/static

COPY ./ /app
WORKDIR /app/src

RUN go test
RUN go build main.go

EXPOSE 80
CMD ./main
