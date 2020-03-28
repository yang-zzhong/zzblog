
+---+

urlid: zzblog-readme

title:ablout zzblog

tags: #zzblog, #design

category: design

overview: zzblog is a personal blog system, it parse a direction to find images, blog markdowns for the web content.

+---+


## INTRO

I have wanted a blog system that on my taste for a long time, I've designed and implemented it as I had a long rest time after quit the last job. it included a parser to parse the blog head from a extended markdown syntax. gathering the categories, tags from the blog head to organize the contents.

the product sample on [ironwell young's home page](https://iiiboo.cn), if the traffic bad, please be patient because of the host in the china mainland. but I will move it to the out side of the wall ofter I get a credit card. get the code on [github](https://github.com/yang-zzong/zzblog)

## INSTALL

### FRONTEND

edit `frontend/src/model.js`

```js
const model = {
    server: 'http://localhost:8080', // edit this line to `server: ''`
};
```

then

```bash
$ cd frontend && npm run build
```

#### REQUIREMENT

* npm
* react

### compile and install backend

```bash
$ go get ./... && cd main && go build && mv main /path/to/bin/
```
U can modify the systemd to manage the process through edit a service file follow below.

```bash
$ vim zzblog.service
```

copy below content to the file

```
[Unit]
Description=zz Blog
Wants=network.target

[Service]
ExecStart=/path/to/zzblog start -c /path/to/config.yml
ExecReload=/bin/kill -HUP $MAINPID
Restart=on-failure
ProtectSystem=true
User=root
Group=root
```

and modify the `/path/to/zzblog` to the abs path of the zzblog, and the abs path of the config.yml

```bash
$ mv zzblog.service /etc/systemd/system/
```

#### REQUIREMENT

* golang
* **rendertron** for the server renderer to help search engine to spide the content

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

## BLOG STRUCTURE

### DIR STRUCTURE

the dir structure is below

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

you can create sub direction in blogs and images, the init program will search all the `/root/blogs` for blogs

### EXTENDED MARKDOWN

I expended the markdown through add a head, all head example below

```markdown

+---+
title: hello world
urlid: hello-world
overview: a hello world sample of extended markdown
tags: #hello world, #extended markdown
cate: nornal cate
image: /api/hello-world.jpg
lang: en
+---+
## title one
content
## title two
content

```
