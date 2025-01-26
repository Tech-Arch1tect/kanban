FROM nginx:alpine

COPY ./client/dist /usr/share/nginx/html

COPY ./client/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
