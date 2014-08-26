docker run --name mysql -e MYSQL_DATABASE=wordpress -e MYSQL_ROOT_PASSWORD=mysecretpassword -d mysql

docker run -p 8081:80 -v /Users/ben/Desktop/wordpress:/var/www/html --name wordpress --link mysql:mysql -e WORDPRESS_DB_HOST=mysql -e WORDPRESS_DB_USER=root -e WORDPRESS_DB_NAME=wordpress -e WORDPRESS_DB_PASSWORD=mysecretpassword -d wordpress