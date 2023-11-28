FROM nginx:latest

WORKDIR /var/html

COPY . .

FROM nginx:latest

EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]

