
+---+

urlid: zzblog-readme

title: zzblog设计

tags: #zzblog, #design

category: design

overview: zzblog是一个个人小型博客系统，包含一个从拓展的markdown中解析元数据的工具，一个博客索引工具。

+---+


## 介绍

因为自身需要，所以想要创建一个博客后端，注意，是后端。目前市面上存在的博客系统是前后端没有分离的产品，那么，当有前端做单页应用需求时则很难满足。基于此。该项目尽提供HTTP API，且这些API仅获取数据。create数据通过提供命令行功能完成。这样，我们可以简化认证流程。

## 命令行工具

1. run
2. post         
3. image

## HTTP API

### 获取博客列表
```
GET /blogs
```

#### request
```
page={page}                 # 可选，第几页, 默认1
page_size={page_size}       # 可选，页大小，默认10
tag={tag}                   # 可选，包含标签
cate={cate}                 # 可选，分类
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
        images: []string,
        viewes: int,
        comments: int,
        thumb_ups: int,
        thumb_downs: int,
        created_at: time_string,
        updated_at: time_string,
    }]
}
```

### 获取博客详情
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
    images: []string,
    viewes: int,
    comments: int,
    thumb_ups: int,
    thumb_downs: int,
    created_at: time_string,
    updated_at: time_string,
    content: markdown_string
}
```

### 获取分类
```
GET /cates
```

#### requset
会下发所有分类，因为分类数量不大

#### response
```
Content-Type: application/json

[]string
```

### 获取所有标签
```
GET /tags
```

#### response
```
Content-Type: application/json

[]string
```

## 命令行接口

### 运行HTTP服务器
```
run [path-prefix=/api] [doc-root=/var/www/zzblog] [with-seo] [with-spa]
```

1. path-prefix 该选项指明是否需要在HTTP API之前加路径。比如 path-prefix=/api，那么请求文章详情的接口就编程了```GET /api/blogs/:url_id```, 该选项的用途主要是在有doc-root且是spa的情况下区分接口和前端页面
2. doc-root 静态页面的根目录。入口文件未 index.html
3. with-seo 启动seo优化的功能，当发布文章时，会同时生成静态的html文件。当检测到spider请求时，会直接将静态文件发给它
4. with-spa 是否以单页应用方式运行服务器

### 创建文章

```
post [with-seo] file.md
```
#### 文章格式

```
<title>title</title>
<urlid>if-it-is-friday</urlid>
<tags>[tag1][tag2][tag3]</tags>
<cate>category</cate>
<overview>overview</overview>

剩下的均为content
```

例子
```
<title>如果今天是星期五</title>
<urlid>if-it-is-friday</urilid>
<tags>[星期五]</tags>
<cate>日常</cate>
<overview>今天是星期五，明天要放周末了，怎么玩呢</overview>

## 简介
今天是星期五了哟
## 第一步
## 第二步
```

### 上传图片
```
image
```
