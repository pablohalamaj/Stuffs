# "ported" by Adam Miller <maxamillion@fedoraproject.org> from
#   https://github.com/fedora-cloud/Fedora-Dockerfiles
#
# Originally written for Fedora-Dockerfiles by
#   Dusty Mabe <dustymabe@gmail.com>

FROM centos:centos7
MAINTAINER The CentOS Project <cloud-ops@centos.org>

# Perform updates
RUN yum -y update; yum clean all

# Install EPEL for owncloud packages
RUN yum -y install httpd ;yum clean all
RUN yum -y install mod_ssl;yum clean all
RUN yum -y install epel-release; yum clean all

# Install Pyd.io Repos
RUN rpm -Uvh http://dl.ajaxplorer.info/repos/pydio-release-1-1.noarch.rpm
RUN yum update -y
RUN yum --disablerepo=pydio-testing install pydio -y
RUN chown -R apache:apache /var/lib/pydio/
#RUN service httpd restart
# Adapto la config de httpd 2.2 a httpd 2.4
RUN sed -i "s/Allow from all/Require all granted/" /etc/httpd/conf.d/pydio.conf
RUN sed -i "s/Order allow,deny/#Order allow,deny/" /etc/httpd/conf.d/pydio.conf
RUN sed -i "s/Deny from all/Require all denied/" /etc/httpd/conf.d/pydio.conf
# Desactivo el output_buffering
RUN sed -i "s/^output_buffering =.*/output_buffering = off/" /etc/php.ini
# Exporto el Lenguaje
RUN export LANG=en_us.UTF-8
###RUN sed -i 's;//define("AJXP_LOCALE", "en_EN.UTF-8");define("AJXP_LOCALE", "en_EN.UTF-8");' /etc/pydio/bootstrap_conf.php

# Expose port 443 and set httpd as our entrypoint
EXPOSE 443
ENTRYPOINT ["/usr/sbin/httpd"]
CMD ["-D", "FOREGROUND"]

