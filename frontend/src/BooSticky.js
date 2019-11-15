import React from 'react';
import {helper} from './helper';

const raisedStyle = {
  position: 'fixed',
  zIndex: 100
};

const style = {
  position: 'static'
};

export default class BooSticky extends React.Component {
  constructor(props) {
    super(props);
    this.scrollTimer = null;
    this.state = {
      style: style,
      raised: false
    }
  }

  render() {
    return (
      <div ref="wrapper">
        <div ref="content" style={this.state.style}>
          {this.props.children}
        </div>
      </div>
    );
  }

  componentDidMount() {
    this._onScroll = this.onScroll.bind(this);
    this._onResize = this.onResize.bind(this);
    window.addEventListener('scroll', this._onScroll);
    window.addEventListener('resize', this._onResize);
  }

  componentWillUnmount() {
    window.removeEventListener('scroll', this._onScroll);
    window.removeEventListener('resize', this._onResize);
  }

  onScroll() {
    clearTimeout(this.scrollTimer);
    this.scrollTimer = setTimeout(() => {
      this.updateSticky();
    }, 1);
  }

  onResize() {
    setTimeout( () => this.updateSticky(), 100);
  }

  updateSticky() {
    if (this.props.disabled) {
      return;
    }
    if (this.props.top !== undefined) {
      this._stickyTop(this.props.top);   
    } else if (this.props.bottom !== undefined) {
      this._stickyBottom(this.props.bottom);
    }
  }

  _stickyTop(top) {
    let wrapper = this.wrapper();
    if (!wrapper) {
      return;
    }
    let node = this.content();
    let raised = this.state.raised;
    this.setState({style: style, raised: false});
    wrapper.style.height = 'auto';
    node.style.width = 'auto';
    if (helper.window_scroll_top() + top > helper.pos_in_doc(node).y) {
      let r = node.getBoundingClientRect();
      if(r.width > 0) {
        node.style.width = r.width + 'px';
      }
      if (r.height > 0) {
        wrapper.style.height = r.height + 'px';
      }
      node.style.top = top + 'px';
      this.setState({style: raisedStyle, raised: true});
    }
    if (raised !== this.state.raised && this.props.onRaised) {
      this.props.onRaised(this.state.raised);
    }
  }

  _stickyBottom (bottom) {
    let wrapper = this.wrapper();
    let node = this.content();
    if (!wrapper) {
      return;
    }
    let raised = this.state.raised;
    this.setState({style: style, raised: false});
    wrapper.style.height = 'auto';
    node.style.width = 'auto';
    let r = node.getBoundingClientRect();
    let top = helper.screen_height() - bottom - r.height;
    if (helper.window_scroll_top() + top < helper.pos_in_doc(node).y) {
      if(r.width > 0) {
        node.style.width = r.width + 'px';
      }
      if (r.height > 0) {
        wrapper.style.height = r.height + 'px';
      }
      node.style.bottom = bottom + 'px';
      this.setState({style: raisedStyle, raised: true});
    }
    if (raised !== this.state.raised && this.props.onRaised) {
      this.props.onRaised(this.state.raised);
    }
  }

  wrapper() {
    return this.refs.wrapper;
  }

  content() {
    return this.refs.content;
  }
}
