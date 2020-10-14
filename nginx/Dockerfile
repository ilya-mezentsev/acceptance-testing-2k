FROM nginx:latest

ARG NGINX_CONF_FILE
COPY ./conf /etc/nginx/at2k
COPY $NGINX_CONF_FILE /etc/nginx/nginx.conf
EXPOSE 80
EXPOSE 443
