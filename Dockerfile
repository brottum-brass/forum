FROM nginx:1.31-alpine3.24

COPY static/html /usr/share/nginx/html/

EXPOSE 80
