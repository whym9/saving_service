#! /bin/sh
sudo docker run -e MYSQL_ROOT_PASSWORD=123 --name=mys -p 127.0.0.1:3307:3306 -d mysql
sudo docker exec -i mys mysql -uroot -p123 < db.sql
sudo docker build -t saver . 
sudo docker run --name=save --env-file .env -p 5005:5005 -d saver
