FROM ubuntu:latest

EXPOSE 8000

WORKDIR /app

ENV DBHOST=localhost DBPORT=5432

ENV DBUSER=root DBPASS=root DBNAME=root


COPY ./main main

CMD ["./main"]
