dockerfiles-centos-OwnCloud
===========================

This repo contains a recipe for making Docker container for Pyd.io on CentOS7. 
The docker file chooses httpd and sqlite for owncloud. These can easily be changed
by changing what rpms get installed. This has been tested with docker 1.0.0 

Check your Docker version

    # docker version

Perform the build

    # docker build -t <yourname>/pyd.io:centos7 .

Check the image out.

    # docker images

Run it:

    # docker run -d -p 443:443 <yourname>/pyd.io:centos7

You should now be able to view the pyd.io setup page by going to https://localhost/pyd.io

NOTE: Ignore the warning about webdav not being set up properly. This is because the SSL 
      certificates aren't set up properly.

