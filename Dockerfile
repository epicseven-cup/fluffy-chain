FROM golang:latest
WORKDIR /fluffychain/
COPY . .
RUN go install
RUN go build 
CMD ["go", "run", "fluffychain"]