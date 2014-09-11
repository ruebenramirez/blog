#!/bin/bash

set +e

WORKING_DIR=`pwd`

docker build -t ruebenramirez/blog.ruebenramirez.com .

# docker push ruebenramirez/blog.ruebenramirez.com

docker rm -f blog.ruebenramirez.com

docker run -d \
    -p 80:80 \
    -v $WORKING_DIR/output/:/var/www/ \
    -v $WORKING_DIR/nginx:/etc/nginx/sites-enabled/ \
    --name blog.ruebenramirez.com \
    ruebenramirez/blog.ruebenramirez.com

