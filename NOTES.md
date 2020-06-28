# How To

## See what ports are opened by firewall

```
firewall-cmd --list-all
```

## Add port to firewall

```
firewall-cmd --zone=public --add-port=8080/tcp --permanent
firewall-cmd --reload
```

## Remove port from firewall

```
firewall-cmd --zone=public --remove-port=8080/tcp
firewall-cmd --runtime-to-permanent
firewall-cmd --reload
```

## Nginx

### Install and enable

```
yum install nginx
systemctl enable nginx.service
systemctl start nginx.service
```

### Configure

create file `/etc/nginx/conf.d/main.conf`

```
server {

    # listen to port 80
    listen [::]:80;
    listen 80;

    # from any request accessing yourdomain.com
    # server_name yourdomain.com;

    # forward it to port 3000
    location / {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_next_upstream error timeout http_502 http_503 http_504;
    }

}
```

### Checking and reload

```
nginx -t
systemctl reload nginx.service
```

PS: You might probably need to: `setsebool -P httpd_can_network_connect 1`

# Install Docker

```
sudo yum install -y yum-utils
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
yum install docker

systemctl enable docker.service
systemctl start docker.service

pip3 install docker-compose
docker-compose up -d
```
