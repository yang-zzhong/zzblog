+---+

urlid: zzblog-readme

title: zzblog - 一个个人博客系统

tags: #zzblog, #design

category: zzblog

overview: zzblog是一个个人博客系统，包含一个拓展的markdown解释器，它能解析一个markdown的头部和内容，然后交给web服务器展示起内容。jj

lang: zh-CN
+---+


## INTRO

因为辞去了上一份工作，因此我有了一段空闲时间，正好借助这段时间设计和实现我一直都想着手做的一个个人博客系统。虽然晚上已经有类似的开源项目，但是其前段的用户体验不满足我的要求，且没办法前后端分离。该项目包含

* 一个markdown拓展的解释器
* react实现的单页应用前端
* 依赖rendertron的服务端呈现帮助爬虫发现该网站内容

我已经将这些代码用于我的个人博客 [薛定谔的猫](https://iiiboo.cn), 如果网络很差，请保持耐心，该服务器在中国大陆. 当我拿到属于我的信用卡之后我会将这个服务迁移到墙外. 从[github](https://github.com/yang-zzong/zzblog)获取其源代码

## 安装 

从github上获取源代码

```
$ git clone https://github.com/yang-zzhong/zzblog && cd zzblog
```

### 前端
编辑`frontend/src/model.js`

```js
const model = {
    server: 'http://localhost:8080', // 将这行代码改为 `server: ''`
};
```

之后

```bash
$ cd frontend && npm run build
```

#### 需求

* npm
* react

### 编译及安装后端

```bash
$ go get ./... && cd main && go build && mv main /path/to/bin/
```
你可以修改根据下面的步骤使用systemd去管理该服务

```bash
$ vim zzblog.service
```

复制下面的内容到zzblog.service

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

修改`/path/to/zzblog`成zzblog的绝对路径, 修改`/path/to/config.yml`成config.yml的绝对路径

```bash
$ mv zzblog.service /etc/systemd/system/
```

#### 需求

* golang
* **rendertron** 一个服务端程序的工具去帮助爬虫发现该网站的内容

## HTTP接口

### 获取博客列表

```
GET /blogs
```

#### 请求
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
GET /blogs/:url_id
```

#### 返回
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

### 获取分类

```
GET /cates
```

#### 返回
```
Content-Type: application/json

[]string
```

### 获取标签

```
GET /tags
```

#### 返回
```
Content-Type: application/json

[]string
```

## 博客结构

### 目录结构

完整的目录结构如下

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

你可以在`/root/blogs`和`/root/images`创建子目录,初始化时程序会遍历`/root/blogs`去发现文章

### 拓展的MARKDOWN

我给markdown添加了个头，来拓展markdown方便定义每篇博客的基础信息

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
