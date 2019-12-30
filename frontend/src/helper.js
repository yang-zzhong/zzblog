import MD5 from 'md5-es';

export const helper = {
  updateKeywords: function(keywords) {
    let meta = document.querySelector('[name=keywords]');
    meta.setAttribute('content', keywords.join(','));
  },
  updateDescription: function(description) {
    let meta = document.querySelector('[name=description]');
    meta.setAttribute('content', description);
  },
  updateTitle: function(title) {
    document.title = title;
  },
  to_show_img: function(img) {
    img.addEventListener('click', e => {
      let id = e.target.getAttribute("data-id");
      if (e.target.hasAttribute('data-id')) {
        window.boo.location.go("/images/" + id.substr(1));
        return;
      }
      if (e.target.hasAttribute('src')) {
        window.boo.location.go('/image?src=' + encodeURIComponent(e.target.getAttribute('src')));
        return;
      }
    })
    img.addEventListener('mouseenter', e => {
      e.target.style.cursor = 'pointer';
    })
    img.addEventListener('mouseout', e => {
      e.target.style.cursor = 'auto';
    });
  },
  rand_str: function(len) {
    return 'a' + Math.random().toString(len + 2).substr(2);
  },
  name_from_mail: function(str) {
    return str.replace(/@.*/i, '');
  },
  is_phone: function(str) {
    return /^[1][3,4,5,7,8][0-9]{9}$/.test(str); 
  },
  pos_in_doc: function(node) {
    var pos = {x: node.offsetLeft, y: node.offsetTop};
    var p = node.offsetParent;
    while( p != null ) {
      pos.x += p.offsetLeft + (p.offsetWidth - p.clientWidth) / 2;
      pos.y += p.offsetTop + (p.offsetHeight - p.clientHeight) / 2;
      p=p.offsetParent;
    }
    return pos;
  },
  window_scroll_top: function() {
    let el = document.scrollingElement || document.documentElement;
    return el.scrollTop;
  },
  screen_height: function() {
    if (document.documentElement.clientHeight) {
      return document.documentElement.clientHeight;
    }
    return document.body.clientHeight;
  },
  doc_height: function() {
    if (document.documentElement.scrollHeight) {
      return document.documentElement.scrollHeight;
    }
    return document.body.scrollHeight;
  },
  screen_width: function() {
    if (document.documentElement.clientWidth) {
      return document.documentElement.clientWidth;
    }
    return document.body.clientWidth;
  },
  index_of_blog: function(node) {

    const insert_index = (root, item) => {
      let len = root.children.length - 1;
      if (root.children.length === 0 || root.weight === item.weight - 1) {
        root.children.push(item);
        return;
      }
      if (root.children[len].weight >= item.weight) {
        root.children.push(item);
        return;
      }
      insert_index(root.children[len], item);
    };

    let root = {weight: 0, children: []};
    node.querySelectorAll('h1,h2,h3,h4').forEach((node) => {
      const id = "a" + MD5.hash(node.innerHTML);
      node.setAttribute("id", id);
      let item = {
        url: '#' + id,
        name: node.textContent,
        opened: true,
        weight: parseInt(node.tagName.substr(1,1)),
        children: []
      };
      if (item.name.trim() === "") {
        return;
      }
      insert_index(root, item);
    });

    return root.children;
  },
  is_android_agent: function() {
    let u = navigator.userAgent;
    return u.indexOf('Android') > -1 || u.indexOf('Adr') > -1;
  },
  is_ios_agent: function() {
    let u = navigator.userAgent;
    return u.indexOf('iPhone') > -1 || u.indexOf('iPad') > -1;
  },
  is_mail: function(str) {
    let re=/^\w+((-\w+)|(\.\w+))*@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z0-9]+$/;
    return re.test(str);
  }
};