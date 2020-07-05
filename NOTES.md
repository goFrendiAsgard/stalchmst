# How To

## See what ports are opened by firewall

```sh
firewall-cmd --list-all
```

## Add port to firewall

```sh
firewall-cmd --zone=public --add-port=8080/tcp --permanent
firewall-cmd --reload
```

## Open masquerade

```sh
firewall-cmd --zone=public --add-masquerade
firewall-cmd --permanent --zone=public --add-masquerade
```

## Remove port from firewall

```sh
firewall-cmd --zone=public --remove-port=8080/tcp
firewall-cmd --runtime-to-permanent
firewall-cmd --reload
```

## No ping

No ping to the outside world. The chances you are missing `sysctl -w net.ipv4.ip_forward=1`

And do not forget to make it permanent by adding the `"net.ipv4.ip_forward=1"` to `/etc/sysctl.conf` (or a file `.conf` in `/etc/sysctl.d/`).

## Nginx

### Install and enable

```sh
yum install nginx
systemctl enable nginx.service
systemctl start nginx.service
```

### Configure

create file `/etc/nginx/conf.d/main.conf`

```nginx
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

```sh
nginx -t
systemctl reload nginx.service
```

PS: You might probably need to: `setsebool -P httpd_can_network_connect 1`

## Install Docker (CentOS)

```sh
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

## Make docker able to access internet

```sh
nslookup google.com # dns server will be shown
```

You can use --net=host in docker run command

```sh
docker run --net=host -it ubuntu
```

Else add dns in config file in /etc/default/docker

```sh
DOCKER_OPTS="--dns 208.67.222.222 --dns 208.67.220.220"
```

for more info refer: [](http://odino.org/cannot-connect-to-the-internet-from-your-docker-containers/)

## Neovim FTW

### Basic setup

```vim
set nocompatible

" enable syntax and plugin
syntax enable
filetype plugin on
```

### Finding files

```vim
set path+=**
set wildmenu
find filename
```

### netrw

```vim
Explore
Sexplore
Vexplore
```

### copy paste to system clipboard

```vim
"*yy
"*p
```

### make clipboard works for wsl

```sh
curl -sLo/tmp/win32yank.zip https://github.com/equalsraf/win32yank/releases/download/v0.0.4/win32yank-x64.zip
unzip -p /tmp/win32yank.zip win32yank.exe > /tmp/win32yank.exe
chmod +x /tmp/win32yank.exe
sudo mv /tmp/win32yank.exe /usr/bin # move win32yank.exe to $PATH
```

### Simplest `~/.config/nvim/init.vim`

```sh
set nocompatible
set wildmenu
syntax enable
filetype plugin on
filetype plugin on

set clipboard+=unnamedplus

call plug#begin()

Plug 'neoclide/coc.nvim', {'branch': 'release'}

call plug#end()
```

### Coc

```vim
CocInstall coc-json coc-tsserver coc-go \
coc-python coc-yaml coc-css coc-html coc-eslint coc-markdownlist \
coc-sh coc-sql coc-vimlsp coc-todolist
```

### vimplug

```sh
sh -c 'curl -fLo "${XDG_DATA_HOME:-$HOME/.local/share}"/nvim/site/autoload/plug.vim --create-dirs \
https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim'

```

