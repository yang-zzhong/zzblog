import model from './model';
import { helper } from './helper';

export const theme = {
  list: [],
  init: function() {
    return model.themes().then(body => {
      theme.list = [];
      for (let i in body) {
        let item = {
          name: body[i].name || i,
          content: body[i]
        };
        theme.list.push(item);
      }
      return new Promise(r => r(theme.list));
    }, (status, msg) => {
      console.error(status, msg);
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
        if (i === '--head-bg-img' && content[i]) {
          let val = model.imageUrl(content[i], helper.screen_width());
          root.style.setProperty(i, 'url(' + val + ')');
          continue;
        }
        root.style.setProperty(i, content[i]);
      }
      localStorage.setItem("theme", name);
    }
  }
};
