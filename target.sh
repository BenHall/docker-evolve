docker build -t docker-node-example github.com/BenHall/docker-node-example
docker run -d --name node-app-1 docker-node-example

echo "upstream node-app-1-3000 { " >> node-app-1-nginx.conf
echo "    server node-app-1:3000; " >> node-app-1-nginx.conf
echo "} " >> node-app-1-nginx.conf
echo "server { " >> node-app-1-nginx.conf
echo "  listen 80 default_server; " >> node-app-1-nginx.conf
echo "  listen [::]:80 default_server ipv6only=on; " >> node-app-1-nginx.conf
echo "  server_name localhost; " >> node-app-1-nginx.conf
echo "  location / { " >> node-app-1-nginx.conf
echo "	  proxy_set_header X-Real-IP \$remote_addr; " >> node-app-1-nginx.conf
echo "    proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for; " >> node-app-1-nginx.conf
echo "    proxy_set_header Host \$http_host; " >> node-app-1-nginx.conf
echo "    proxy_set_header X-NginX-Proxy true; " >> node-app-1-nginx.conf
echo "    proxy_pass http://node-app-1-3000; " >> node-app-1-nginx.conf
echo "    proxy_redirect off; " >> node-app-1-nginx.conf
echo "  } " >> node-app-1-nginx.conf
echo "} " >> node-app-1-nginx.conf

docker run -d -p 8080:80 --name nginx_root --link node-app-1:node-app-1 -v /Users/ben/Desktop/nginx/www:/data -v /Users/ben/Desktop/nginx/sites-enabled:/etc/nginx/sites-enabled -v /Users/ben/Desktop/nginx/logs:/var/log/nginx dockerfile/nginx
