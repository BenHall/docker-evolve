docker build -t docker-node-example github.com/BenHall/docker-node-example
docker run -p 46169:3000 -d docker-node-example

# Inspect docker instance for IP and ideally the random port
# docker-node-example.domain

cat "upstream node-app-1-49160 { " >> node-app-1-nginx.conf
cat "    server 192.168.59.103:49160; " >> node-app-1-nginx.conf
cat "} " >> node-app-1-nginx.conf
cat "server { " >> node-app-1-nginx.conf
cat "  listen 80 default_server; " >> node-app-1-nginx.conf
cat "  listen [::]:80 default_server ipv6only=on; " >> node-app-1-nginx.conf
cat "  server_name localhost; " >> node-app-1-nginx.conf
cat "  location / { " >> node-app-1-nginx.conf
cat "	proxy_set_header X-Real-IP $remote_addr; " >> node-app-1-nginx.conf
cat "    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; " >> node-app-1-nginx.conf
cat "    proxy_set_header Host $http_host; " >> node-app-1-nginx.conf
cat "    proxy_set_header X-NginX-Proxy true; " >> node-app-1-nginx.conf
cat "    proxy_pass http://node-app-1-49160; " >> node-app-1-nginx.conf
cat "    proxy_redirect off; " >> node-app-1-nginx.conf
cat "  } " >> node-app-1-nginx.conf
cat "} " >> node-app-1-nginx.conf




docker run -d -p 8080:80 --name nginx_root --link docker-node-example:docker-node-example -v /Users/ben/Desktop/nginx/www:/data -v /Users/ben/Desktop/nginx/sites-enabled:/etc/nginx/sites-enabled -v /Users/ben/Desktop/nginx/logs:/var/log/nginx dockerfile/nginx
