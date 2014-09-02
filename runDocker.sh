#!/bin/bash

docker run -d \
    -p 80:80 \
    -v ./nginx /etc/nginx/sites-enabled/ \
    -v ./content /var/www/ \
    --name blog.ruebenramirez.com \
    nginx 
