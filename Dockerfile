FROM golang:1.18 
LABEL maintainer ="obintou-bballaki-skaridja"\
      description ="projet ascii-art-web dockerise"

WORKDIR /go/src/projet
COPY . /go/src/projet

RUN go build -o main .
CMD ["./main"]