FROM httpd:latest

COPY ./index.html /usr/local/apache2/htdocs/index.html
COPY ./httpd.conf /usr/local/apache2/conf/httpd.conf

