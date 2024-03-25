FROM golang:1.16.15-alpine3.15
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go build -o mentat .
CMD ["./mentat"]