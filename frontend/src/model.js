
import 'boo-route';
import {localizer} from './localizer';

const model = {
  // server: 'http://localhost:8080'
  server: ''
};

model.clearOldBlogQuery = function() {
  model.blogs_params = null;
}

model.themes = function() {
  return model.fetch(model.url('/api/theme'));
};

model.url = function(path, params) {
  let u = model.server + path;
  let query = "";
  let arr = [];
  for (let i in params) {
    if (i && params[i]) {
      arr.push(i + "=" + params[i]);
    }
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
  const l = localizer.lang();
  if (model._userInfo && model._userInfo[l]) {
    return new Promise(r => r(model._userInfo[l]));
  }
  return new Promise((r, reject) => {
    model.fetch(model.url('/api/author', {lang: l})).then(info => {
      model._userInfo = model._userInfo || {};
      model._userInfo[l] = info;
      r(info);
    }, (status, msg) => {
      reject(status, msg);
    });
  });
};

model.nextBlogs = function() {
  let params = model.blogs_params;
  params.page += 1;
  params.lang = localizer.lang();
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
  params = params || {};
  params['lang'] = localizer.lang();
  let url = model.url('/api/blogs', params);
  return model.fetch(url);
};

model.blog = function() {
  const ctx = window.boo.location.context;
  const path = '/api/blogs/' + ctx.path_params.url_id;
  return model.fetch(model.url(path, {lang: localizer.lang()}));
}

model.imageUrl = function(img, w, h) {
  if (!img || img.search('http') === 0) {
    return img;
  }
  if (img.search('/api') < 0) {
    img = "/api/images/" + img;
  }

  return model.url(img, {w: w, h: h});
}

model.cates = function() {
  return model.fetch(model.url('/api/cates', {lang: localizer.lang()}));
}

model.tags = function() {
  return model.fetch(model.url('/api/tags', {lang: localizer.lang()}));
}

export default model;
