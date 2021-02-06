FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD=pwwd

COPY conf/my.cnf /etc/mysql/conf.d/my.cnf
COPY world.sql.gz docker-entrypoint-initdb.d
