FROM golang:1.17.13-alpine3.16

ARG review_svc_host
ENV REVIEW_SVC_HOST $review_svc_host

RUN mkdir -p /home/app/product-details
COPY . /home/app/product-details
WORKDIR /home/app/product-details

RUN go mod download
RUN go build -o  /product-details
CMD [ "/product-details" ]