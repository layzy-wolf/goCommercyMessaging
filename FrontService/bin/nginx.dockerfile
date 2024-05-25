FROM nginx:1.20.2

RUN rm /etc/nginx/nginx.conf

COPY ./config/nginx.conf /etc/nginx/nginx.conf