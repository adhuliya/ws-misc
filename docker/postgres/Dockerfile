FROM postgres:13

WORKDIR /
COPY setup_postgres.sh .

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-11.1`` package when it was ``apt-get installed``
COPY postgresql.conf /var/lib/postgresql
RUN chown postgres:postgres /var/lib/postgresql/postgresql.conf
RUN chmod -vR g-w,o-rwx /var/lib/postgresql/data
RUN ls -l /var/lib/postgresql

USER postgres
RUN postgres -c config_file=/var/lib/postgresql/postgresql.conf
RUN psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
RUN createdb -O docker docker

