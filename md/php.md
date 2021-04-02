#### php

##### 1.下载 php源码

````shell
wget https://www.php.net/distributions/php-7.4.16.tar.gz
````

````shell
tar -zxvf php-7.4.16.tar.gz
````

##### 2.编译安装

````shell
yum install -y openssl-devel libxml2-devel bzip2-devel libcurl-devel libjpeg-devel libpng-devel freetype-devel libmcrypt-devel recode-devel libicu-devel libzip-devel sqlite-devel oniguruma-devel gcc autoconf automake libtool libxslt-devel
````

````sh
./configure --prefix=/usr/local/soft/php7 --with-curl --with-freetype-dir --with-gd --with-gettext --with-iconv-dir --with-kerberos --with-libdir --with-libxml-dir --with-mysqli --with-openssl --with-pcre-regex --with-pdo-mysql --with-pdo-sqlite --with-pear --with-png-dir --with-xmlrpc --with-xsl --with-zlib --enable-fpm --enable-bcmath -enable-inline-optimization --enable-gd-native-ttf --enable-mbregex --enable-mbstring --enable-opcache --enable-pcntl --enable-shmop --enable-soap --enable-sockets --enable-sysvsem --enable-xml --enable-zip --enable-pcntl --with-curl --with-fpm-user --enable-ftp --enable-session --enable-xml 
````

````shell
make && make install
````



##### 3.环境变量

````shell
vim /etc/profile.d/php.sh
````

````shell
export PHP_HOME=/usr/local/soft/php7
export PATH=$PATH:$PHP_HOME/bin
export PATH=$PATH:$PHP_HOME/sbin
````

````shell
source /etc/profile
````

##### 4.修改配置

````shell
cp /usr/local/src/php/php-7.4.16/php.ini-development  /usr/local/soft/php7/lib/php.ini
````

##### 5.配置php-fpm

````shell
cp php-fpm.conf.default /usr/local/soft/php7/etc/php-fpm.conf
````

````shell
cp /usr/local/soft/php7/etc/php-fpm.d/www.conf.default /usr/local/soft/php7/etc/php-fpm.d/www.conf
````

````shell
php-fpm
netstat -anp | grep php-fpm
````

 ##### 6.安装nginx

````shell
wget http://nginx.org/download/nginx-1.18.0.tar.gz
````



````shell
./configure  --prefix=/usr/local/soft/nginx\
             --with-http_stub_status_module\
             --with-http_ssl_module\
             --with-http_gzip_static_module\
````

````shell
make && make install
````

````shell
vim /etc/profile.d/nginx
````

````shell
export NGINX_HOME=/usr/local/soft/nginx
export PATH=$PATH:$NGINX_HOME/sbin
````

````shell
source /etc/profile
````

加入systemctl  service

````shell
vim /etc/init.d/nginx
````

````shell
#!/bin/bash
#Startup script for the nginx Web Server
#chkconfig: 2345 85 15
nginx=/usr/local/soft/nginx/sbin/nginx
conf=/usr/local/soft/nginx/conf/nginx.conf
case $1 in
start)
echo -n "Starting Nginx"
$nginx -c $conf
echo " done."
;;
stop)
echo -n "Stopping Nginx"
killall -9 nginx
echo " done."
;;
test)
$nginx -t -c $conf
echo "Success."
;;
reload)
echo -n "Reloading Nginx"
ps auxww | grep nginx | grep master | awk '{print $2}' | xargs kill -HUP
echo " done."
;;
restart)
$nginx -s reload
echo "reload done."
;;
*)
echo "Usage: $0 {start|restart|reload|stop|test|show}"
;;
esac

````

````
chmod 755 /etc/init.d/nginx
````

````
chkconfig --add nginx
````

````
systemctl status nginx
````







##### 配置nginx 支持 php

注释掉nginx配置文件php-fpm