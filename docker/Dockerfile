FROM golang

COPY . /go/src/api

WORKDIR /go/src/api

RUN go get -d -v

RUN go install -v

EXPOSE 80

CMD ["api"]
