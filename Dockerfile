FROM node:latest

postgresql://postgres:postgres@localhost:5432/hoge_db

mysql://user:password@localhost:3306/db

    image: mysql:8.0
    platform: linux/amd64
    ports:
      - 3306:3306
    environment:
      MYSQL_ROOT_PASSWORD: mysql
      MYSQL_DATABASE: db
      MYSQL_USER: user
      MYSQL_PASSWORD: password