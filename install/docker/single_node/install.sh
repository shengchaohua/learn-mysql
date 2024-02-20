# docker hub: https://hub.docker.com/r/mysql/mysql-server
# source: github.com/mysql/mysql-docker/blob/mysql-server/5.7/Dockerfile
# setting can be found at: dev.mysql.com/doc/refman/5.7/en/docker-mysql-more-topics.html#docker-persisting-data-configuration

# use 5.7.41
docker pull mysql/mysql-server:5.7.41
docker run -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD='true' -e MYSQL_ROOT_PASSWORD='root' -e MYSQL_ROOT_HOST=% \
    --name="mysql_single_node" -d mysql/mysql-server:5.7.41 \
    --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
