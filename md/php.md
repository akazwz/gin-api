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
# chkconfig:   - 85 15

# description:  Nginx is an HTTP(S) server, HTTP(S) reverse \
#               proxy and IMAP/POP3 proxy server

# processname: nginx
# config:      /usr/local/nginx/conf/nginx.conf
# pidfile:     /usr/local/nginx/logs/nginx.pid

# Source function library.

. /etc/rc.d/init.d/functions

# Source networking configuration.

. /etc/sysconfig/network

# Check that networking is up.

[ "$NETWORKING" = "no" ] && exit 0

nginx="/usr/local/soft/nginx/sbin/nginx"

prog=(basename(basenamenginx)

NGINX_CONF_FILE="/usr/local/soft/nginx/conf/nginx.conf"

lockfile=/var/lock/subsys/nginx

start() {

[ -x $nginx ] || exit 5

[ -f $NGINX_CONF_FILE ] || exit 6

echo -n "Starting"Startingprog: "

daemon nginx -cnginx−cNGINX_CONF_FILE

retval=$?

echo

[ $retval -eq 0 ] && touch $lockfile

return $retval

}

stop() {

echo -n "Stopping"Stoppingprog: "

killproc $prog -QUIT

retval=$?

echo

[ $retval -eq 0 ] && rm -f $lockfile

return $retval

}

restart() {

configtest || return $?

stop

start

}

reload() {

configtest || return $?

echo -n "Reloading"Reloadingprog: "

killproc $nginx -HUP

RETVAL=$?

echo

}

force_reload() {

restart

}

configtest() {

nginx -t -cnginx−t−cNGINX_CONF_FILE

}

rh_status() {

status $prog

}

rh_status_q() {

rh_status >/dev/null 2>&1

}

case "$1" in

start)

rh_status_q && exit 0
        $1
        ;;

stop)

rh_status_q || exit 0
        $1
        ;;

restart|configtest)
        $1
        ;;

reload)
        rh_status_q || exit 7
        $1
        ;;

force-reload)
        force_reload
        ;;
    status)
        rh_status
        ;;

condrestart|try-restart)

rh_status_q || exit 0
            ;;

*)

echo "Usage:"Usage:0 {start|stop|status|restart|condrestart|try-restart|reload|force-reload|configtest}"
        exit 2

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