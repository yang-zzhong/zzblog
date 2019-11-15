import 'boo-route';

const model = {
  server: 'http://localhost:8080'
};

model.blogs = function() {
  // const ctx = window.boo.location.context;
  return fetch(model.server + '/api/blogs').then(res => {
    if (res.status === 200) {
      return res.json();
    }
  }).then(body => {
    window.dispatchEvent(new CustomEvent("list-ready", {detail: body}));
    return new Promise(r => r());
  });
}; 

model.blog = function() {
  const ctx = window.boo.location.context;
  console.log(ctx);
  return fetch(model.server + '/api/blogs/' + ctx.path_params.url_id).then(res => {
    if (res.status === 200) {
      return res.json();
    }
  }).then(body => {
    console.log(body);
    window.dispatchEvent(new CustomEvent("blog-ready", {detail: body}));
    return new Promise(r => r());
  });
}

export default model;
