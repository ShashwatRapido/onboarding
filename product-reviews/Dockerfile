FROM golang:1.17.13-alpine3.16

RUN mkdir -p /home/app/product-reviews

COPY . /home/app/product-reviews
WORKDIR /home/app/product-reviews
RUN go mod download
RUN go build -o /product-reviews
CMD [ "/product-reviews" ]