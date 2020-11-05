import React from 'react';
import { helper } from './helper';
import { animation } from './animation';
import SweetScroll from 'sweet-scroll';

export default class Page extends React.Component {

  constructor(props) {
    super(props);
    this.scrollTop = 0;
    this.show = false;
    this.pc = React.createRef();
    window.addEventListener('resize', e => {
      if (!this.show) {
        return;
      }
      this.scrollTop = helper.window_scroll_top();
    });
    window.addEventListener('scroll', e => {
      if (!this.show) {
        return;
      }
      this.scrollTop = helper.window_scroll_top();
    });
    window.addEventListener('lang-changed', e => this.setState({}));
  }

  enter(state) {
    state = state || {state: 200};
    return new Promise(r => r(state));
  }

  leave() {
    this.show = false;
    return new Promise(r => r());
  }

  entryAnimation() {
    return new Promise(r => r());
  }

  exitAnimation() {
    return new Promise(r => r());
  }

  findScrollNode(hash) {
    if (hash && this.pc.current) {
      return this.pc.current.querySelector(hash);
    }
  }

  scrollMarginTop() {
    return 0;
  }

  scrollTo() {
    this.show = true;
    return new Promise(r => {
      let s = new SweetScroll();
      let node = this.findScrollNode(window.location.hash);
      if (node) {
        s.toElement(node, {duration: 1000});
      } else if (this._url !== window.location.href) {
        s.to(0, {});
        this._url = window.location.href;
      } else {
        s.to(this.scrollTop, {duration: 0});
      }
      setTimeout(() => {
        r();
      }, 5);
    });
  }

  isVisible(node) {
    return helper.is_visible(node);
  }

  playAnimation(conf) {
    let wait = 0;
    if (!conf) {
      return new Promise(r => r());
    }
    conf.forEach(i => {
      animation.play(i);
      wait = i.time > wait ? i.time : wait;
    });
    return new Promise(r => {
      setTimeout(() => { 
        r();
      }, wait);
    });
  }
}
