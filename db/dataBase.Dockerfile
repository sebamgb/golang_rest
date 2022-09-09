ARG db_version=10.3

FROM postgres:${db_version}-alpine

COPY ["up.sql", "/docker-entrypoint-initdb.d/1.sql"]

CMD [ "postgres" ]