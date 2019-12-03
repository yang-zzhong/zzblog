
export const theme = {
  list: [],
  init: function() {
    return fetch('/theme.json').then(res => {
      if (res.status === 200) {
        return res.json().then(body => {
          theme.list = [];
          for (let i in body) {
            let item = {
              name: i,
              content: body[i]
            };
            theme.list.push(item);
          }
          return new Promise(r => r(theme.list));
        });
      }
      return new Promise(r => r([]));
    });
  },
  guess: function() {
    return localStorage.getItem('theme');
  },
  use: function(name) {
    let content = false
    for (let i = 0; i < theme.list.length; ++i) {
      if (theme.list[i].name === name) {
        content = theme.list[i].content;
        break;
      }
    }
    if (content) {
      let root = document.documentElement;
      for (let i in content) {
        root.style.setProperty(i, content[i]);
      }
      localStorage.setItem("theme", name);
    }
  }
};
