FROM nginx

WORKDIR /usr/share/nginx/html

COPY ./dist .

COPY ./deploy/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
EXPOSE 443