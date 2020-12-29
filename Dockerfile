FROM golang:latest

WORKDIR /app

COPY . .

ENV ADMIN_USER=admin
ENV ADMIN_PASS=demo
ENV SIGNING_KEY=veryverysecretkey

RUN go install

EXPOSE 8080

ENTRYPOINT ["go-rest-api"]

CMD ["start","-a"]

