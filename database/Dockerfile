FROM postgres:10.4

ENV POSTGRES_DB giggle
ENV POSTGRES_USER giggle
ENV POSTGRES_PASSWORD 123456

COPY ./up.sql /docker-entrypoint-initdb.d/1.sql

EXPOSE 5432
CMD ["postgres"]
