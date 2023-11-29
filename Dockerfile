FROM nginx:latest

WORKDIR /var/html


COPY . /var/html/

EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]

