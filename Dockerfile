FROM dockerfile/nginx

# update server name hash bucket size
RUN sed -i 's/# server_names_hash_bucket_size/server_names_hash_bucket_size/g' /etc/nginx/nginx.conf

# Define mountable directories.
VOLUME ["/etc/nginx/sites-enabled", "/etc/nginx/conf.d", "/var/log/nginx"]

# Define working directory.
WORKDIR /etc/nginx

cmd ["nginx"]

# Expose ports.
EXPOSE 80
EXPOSE 443
