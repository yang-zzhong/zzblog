import 'boo-route';

const model = {
  server: 'http://localhost:8080'
};

model.url = function(path, params) {
  let u = model.server + path;
  let query = "";
  let arr = [];
  for (let i in params) {
    arr.push(i + "=" + params[i]);
  }
  for (let i = 0; i < arr.length; ++i) {
    query += arr[i];
    if (i < arr.length - 1) {
      query += "&";
    }
  }
  if (query !== "") {
    return u + "?" + query;
  }
  return u;
};

model.queryBlogParams = function() {
  const ctx = window.boo.location.context;
  let params = {};
  let param_changed = false;
  if (!model.blogs_params) {
    param_changed = true;
    model.blogs_params = {
      tag: '',
      page: 1
    };
  }
  let cate = '';
  let tag = '';
  let page = 1;
  if (ctx.path_params.catename) {
    cate = ctx.path_params.catename;
  }
  if (ctx.path_params.tagname) {
    tag = ctx.path_params.tagname;
  }
  if (ctx.query_params.page) {
    page = ctx.query_params.page;
  }
  if (model.blogs_params.cate !== cate) {
    model.blogs_params.cate = cate;
    param_changed = true;
    if (cate !== '') {
      params['cate'] = cate;
    }
  }
  if (model.blogs_params.tag !== tag) {
    model.blogs_params.tag = tag;
    param_changed = true;
    if (tag !== '') {
      params["tag"] = tag;
    }
  }
  if (model.blogs_params.page !== page) {
    model.blogs_params.page = params["page"] = page;
    param_changed = true;
  }
  if (!param_changed) {
    return false;
  }
  return params;
}

model.userInfo = function() {
  if (model._userInfo) {
    return new Promise(r => r(model._userInfo));
  }
  return new Promise((r, reject) => {
    model.fetch(model.url('/api/author')).then(info => {
      model._userInfo = info;
      r(info);
    }, (status, msg) => {
      reject(status, msg);
    });
  });
};

model.nextBlogs = function() {
  let params = model.blogs_params;
  // params.page += 1;
  return model.queryBlogs(params);
};

model.fetch = function(url, opt) {
  return new Promise((r, reject) => {
    fetch(url, opt).then(res => {
      if (res.status === 200) {
        res.json().then(body => {
          r(body);
        });
        return;
      }
      res.text().then(msg => reject(res.status, msg));
    });
  });
};

model.queryBlogs = function(params) {
  let url = model.url('/api/blogs', params);
  return model.fetch(url);
};

model.blog = function() {
  const ctx = window.boo.location.context;
  const path = '/api/blogs/' + ctx.path_params.url_id;
  return model.fetch(model.url(path));
}

model.cates = function() {
  return model.fetch(model.url('/api/cates'));
}

model.tags = function() {
  return model.fetch(model.url('/api/tags'));
}

export default model;
