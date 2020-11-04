<!-- +
urlid: zzblog-readme
title:ablout zzblog
tags: #zzblog, #design
category: zzblog
overview: zzblog is a personal blog system, it parse a direction to find images, blog markdowns for the web content.
lang: en
created_at: 2020-11-05 21:23:00
updated_at: 2020-11-05 21:23:00
+ -->

## INTRO

I have wanted a blog system on my taste for a long time, I've designed and implemented it as I had a long rest time after quiting the last job. it included a parser to parse the blog head from a extended markdown syntax. gathering the categories, tags from the blog head to organize the contents.

the product sample on [ironwell young](https://ironwell-young.me),  get the code on [yang-zzhong/zzblog](https://github.com/yang-zzhong/zzblog)

## extended markdown syntax

Markdown is more readable and more writable than html. but it lacks a head to describe the content of the file. So I extended it thru adding a head to contain the meta information of the content. the format likes below.

```
<!-- +
urlid: zzblog-readme
title:ablout zzblog
tags: #zzblog, #design
category: zzblog
overview: zzblog is a personal blog system, it parse a direction to find images, blog markdowns for the web content.
lang: en
created_at: 2020-11-05 21:23:00
updated_at: 2020-11-05 21:23:00
+ -->

```

You can find the parser at [yang-zzhong/md](https://github.com/yang-zzhong/md)。Maybe you notice that I wrap the meta info in `<!-- + + -->`, yes, that includes a comment syntax of html and then the head will auto hide when show the content as a html file. your markdown file content just likes below.

```
<!-- +
urlid: zzblog-readme
title:ablout zzblog
tags: #zzblog, #design
category: zzblog
overview: zzblog is a personal blog system, it parse a direction to find images, blog markdowns for the web content.
lang: en
created_at: 2020-11-05 21:23:00
updated_at: 2020-11-05 21:23:00
+ -->

## title 1

content 1

## title 2

content 2

```

## DIR STRUCTURE

your blog directer structure may like below

```

root --
     |--- blogs
     |     | ----- about.md
     |     | ----- about-zh-cn.md
     |
     \--- images
           | ----- about.gif
           | ----- logo.jpg
        
```

the `root/blogs` dir contains the blog markdown file which describe before. `root/images` dir contains the images that blogs used. blog can use the following method to access the image in the content of the blog

```

![logo](/api/images/logo.jpg)

```

## INSTALL

I use `rendertron` to implement the server renderer for search engine. so learn how to use it. you could find some useful information from [GoogleChrome/rendertron](https://github.com/GoogleChrome/rendertron)

I assume you will put everything you need in dir `/production`. follow following steps to install zzblog service.

### GET THE CODE

```
$ cd /producton && git clone https://github.com/yang-zzhong/zzblog
```

### COMPILE SERVER PROGRAM

```
$ cd zzblog/main/ && go build && mv main ../../zzblog && cd ../.. && chmod 744 zzblog
```

### CONFIG
```
$ cp zzblog/config.yml . && vim config.yml
```

copy the following code to the `config.yml` file

```
port: 8080
allow_cors: false
doc_root: /production/frontend/
log_path: /production/log/
domain: https://your-domain
root: /production/blog/
render_cache_dir: /tmp/
renderer: http://127.0.0.1:3000
bots:
   - Googlebot
   - Baiduspider
   - Slurp
   - Yahoo
   - iaskspider
   - Sogou
   - YodaoBot
   - msnbot
   - 360Spider
```

### USE SYSTEMD

```
$ touch zzblog.service
```

then copy the following content to zzblog.service

```
[Unit]
Description=zz Blog
Wants=network.target

[Service]
ExecStart=/production/zzblog start -c /production/config.yml
ExecReload=/bin/kill -HUP $MAINPID
Restart=on-failure
ProtectSystem=true
User=root
Group=root

```

let's use the systemctl to manage the zzblog servie

```
$ systemctl start zzblog
```

if you don't know how to use systemd, please google it.

### FRONTEND

edit `frontend/src/model.js`

```js
const model = {
    server: 'http://localhost:8080', // edit this line to `server: ''`
};
```

then

```bash
$ cd zzblog/frontend/ && npm run-script build && cp -rf build/ ../../frontend
```

### RENDERTRON

please read [GoogleChrome/rendertron](https://github.com/GoogleChrome/rendertron). after instalation, edit the config file, put your rendertron url to renderer line.

### ENABLE HTTPS

Enable https thru `nginx`, config the `nginx` like below

```

server {

    listen 443 ssl;
    server_name your-domain;

    ssl_certificate     /path/to/fullchain.pem;
    ssl_certificate_key /path/to/privkey.pem;
    ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers         HIGH:!aNULL:!MD5;

    access_log /path/to/access.log;
    error_log /path/to/error.log;

    # enable gzip
    gzip on;
    # gzip_types text/plain application/xml application/x-javascript font/woff image/*;
    gzip_types *;

    location / {
        proxy_pass http://127.0.0.1:8080;
    }
}

# redirect http to https

server {
    listen 80 default_server;

    server_name ironwell-young.me;

    return 301 https://$host$request_uri;
}

```

## HTTP API

### get the blogs
```
GET /blogs
```

#### request
```
page={page}                 # optional
page_size={page_size}       # optional
tag={tag}                   # optional
cate={cate}                 # optional
```

#### response
```
Content-Type: application/json
{
    total: int,
    data: [{
        url_id: string,         
        title: string,
        overview: string,
        tags: string[],
        cate: string,
        lang: string,
        image: string,
        created_at: time_string,
        updated_at: time_string,
    }]
}
```

### blog detail
```
GET /blogs/:url_id
```

#### response
```
Content-Type: application/json
{
    url_id: string,         
    title: string,
    overview: string,
    image: string,
    tags: string[],
    cate: string,
    lang: string,
    created_at: time_string,
    updated_at: time_string,
    content: markdown_string
}
```

### get categories
```
GET /cates
```

#### requset

#### response
```
Content-Type: application/json

[]string
```

### get tags
```
GET /tags
```

#### response
```
Content-Type: application/json

[]string
```
