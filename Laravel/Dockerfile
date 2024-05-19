# Stage 1: PHP-FPM
FROM php:8.3-fpm AS php

WORKDIR /var/www/html

RUN apt-get update && \
    apt-get install -y \
        libpng-dev \
        libjpeg-dev \
        libfreetype6-dev \
        zip \
        unzip \
        git && \
    docker-php-ext-configure gd --with-freetype --with-jpeg && \
    docker-php-ext-install gd pdo pdo_mysql

# Install Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

COPY . .

RUN composer install

RUN php artisan key:generate

EXPOSE 9000

CMD ["php-fpm"]

# Stage 2: Nginx
FROM nginx:stable-alpine AS nginx

COPY --from=php /var/www/html /var/www/html

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8000

CMD ["nginx", "-g", "daemon off;"]
