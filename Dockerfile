# Use la imagen de mysql
FROM mysql:8.0

# Define las variables de entorno para la configuración de MySQL
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=energy
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password

# Habilitar la opción local-infile para MySQL
RUN echo '[mysqld]\nlocal-infile = 1' > /etc/mysql/conf.d/local-infile.cnf

# Copia el archivo de migración y CSV a la carpeta docker-entrypoint-initdb.d en el contenedor
COPY ./database/migration.sql /docker-entrypoint-initdb.d/migration.sql
COPY ./database/test_bia.csv /docker-entrypoint-initdb.d/test_bia.csv

EXPOSE 3306
