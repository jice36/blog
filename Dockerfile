FROM postgres
ENV POSTGRES_PASSWORD 1234
ENV POSTGRES_DB blog
COPY blog.sql /docker-entrypoint-initdb.d/