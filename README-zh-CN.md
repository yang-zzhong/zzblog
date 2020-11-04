<!-- +
urlid: zzblog-readme
title: zzblog - 一个个人博客系统
tags: #zzblog, #design
category: 工具
overview: zzblog是一个个人博客系统，包含一个拓展的markdown解释器，它能解析一个markdown的头部和内容，然后交给web服务器展示起内容。
lang: zh-CN
created_at: 2020-11-05 21:23:00
updated_at: 2020-11-05 21:23:00
+ -->


## 介绍

因为辞去了上一份工作，因此我有了一段空闲时间，正好借助这段时间设计和实现我一直都想着手做的一个个人博客系统。虽然网上已经有类似的开源项目，但是其前端的用户体验不满足我的要求，且没办法前后端分离。正好借由这段时间开发一个符合自己心理预期的项目。该项目完成的功能包括。

* 一个支持markdown拓展的解释程序，该程序为markdown文件定义一个头来携带关于该markdown文件内容的描述
* 一个以HTTP为服务端的服务器程序，该程序解析一系列的markdown文件，生成文章列表，标签，分类等信息，再通过HTTP/HTTPS协议提供给前端
* 一套前端代码用于优雅的展示内容
* 良好的支持搜索引擎优化
* 支持多语言，这种多语言不仅仅是程序支持，而且也支持内容根据不同的语言进行切换。如果我们选择英文，那么整个网站都会呈现英文的内容，包括文章内容也要能够根据语言进行切换

目前这些代码已经用于我的个人博客 [薛定谔的猫](https://ironwell-young.me), 其源代码可以在[github](https://github.com/yang-zzhong/zzblog)找到

## 拓展的MARKDOWN协议

markdown比html更易读，可以手工编写，但是markdown目前不能携带一些基本的描述信息。比如这篇markdown的标题，类别，标签等等。如果我们用markdown文件之外的东西来记录这些信息，事情就显得相当麻烦了。因此，我将markdown格式定义了一个头，让他能携带一些元信息。大概样子如下

```

<!-- +
urlid: zzblog-readme
title: zzblog - 一个个人博客系统
tags: #zzblog, #design
category: 工具
overview: zzblog是一个个人博客系统，包含一个拓展的markdown解释器，它能解析一个markdown的头部和内容，然后交给web服务器展示起内容。
lang: zh-CN
created_at: 2020-11-05 21:23:00
updated_at: 2020-11-05 21:23:00
+ -->

```

该头的解释程序可以在[yang-zzhong/md](https://github.com/yang-zzhong/md)找到。可能你注意到了，头部是放在 `<!-- + + -->` 这里面的，我使用了`html`的注释来作为头部标记的一部分。这样子我们不需要额外的隐藏，当该文件解析为http格式展示时就会自动隐藏头部信息。因此，当我们使用该项目发布博客的时候，其每个博客的md文件大概长下面这个样子。

```
<!-- +
urlid: zzblog-readme
title: zzblog - 一个个人博客系统
tags: #zzblog, #design
category: 工具
overview: zzblog是一个个人博客系统，包含一个拓展的markdown解释器，它能解析一个markdown的头部和内容，然后交给web服务器展示起内容。
lang: zh-CN
created_at: 2020-11-05 21:23:00
updated_at: 2020-11-05 21:23:00
+ -->

## 标题二

内容二

## 标题三

内容三

```

## 文件结构

你的博客内容需要按如下的目录结构组织

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

其中`root/blogs`下面放md文件，也就是内容就是上面提到的。`root/images`放图片文件。在博客中访问图片的方式如下

```

![logo](/api/images/logo.jpg)

```

## 安装 

搜索引擎优化用的rendertron。因此我们需要使用到rendertron。怎么使用可以参见[GoogleChrome/rendertron](https://github.com/GoogleChrome/rendertron), 安装步骤如下。
我们假定您会将整个环境放在 `/production` 目录下。

### 获取源代码

从github上获取源代码

```
$ cd /producton && git clone https://github.com/yang-zzhong/zzblog
```

### 编译服务端程序

```
$ cd zzblog/main/ && go build && mv main ../../zzblog && cd ../.. && chmod 744 zzblog
```

### 配置

```
$ cp zzblog/config.yml . && vim config.yml
```

其配置项包括

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

### systemd管理服务

```
$ touch zzblog.service
```

然后copy下面的内容到文件zzblog.service

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

然后你就可以用下面的命令启动你的zzblog服务了

```
$ systemctl start zzblog
```

如果不清楚怎么使用systemd管理服务，请自行Google。

### 前端

编辑`zzblog/frontend/src/model.js`

```js
const model = {
    server: 'http://localhost:8080', // 将这行代码改为 `server: ''`
};
```

之后

```bash
$ cd zzblog/frontend/ && npm run-script build && cp -rf build/ ../../frontend
```

### 安装rendertron

请自行查阅 [GoogleChrome/rendertron](https://github.com/GoogleChrome/rendertron), 安装之后修改配置文件的renderer配置为你的rendertron服务的url地址即可

### 启用https

https需要借助nginx代理来实现。配置如下

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

# 将http请求自动重定向到https的请求
server {
    listen 80 default_server;

    server_name ironwell-young.me;

    return 301 https://$host$request_uri;
}

```

## 服务端HTTP请求

### 博客列表

```

GET /api/blogs

```

#### 请求参数

```
page={page}                 # optional
page_size={page_size}       # optional
tag={tag}                   # optional
cate={cate}                 # optional
```

#### 返回

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

### 博客详情

```

GET /api/blogs/:urlid

```

#### 返回
```
Content-Type: application/json
{
    urlid: string,         
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

### 获取分类

```
GET /api/cates
```

#### 返回

```
Content-Type: application/json

[]string

```

### 获取标签

```

GET /api/tags

```

#### 返回

```

Content-Type: application/json

[]stringko

```
### 获取作者信息

```

GET api/author

```
