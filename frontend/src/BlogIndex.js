
import React from 'react';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import './blogindex.css';

export default class BooIndex extends React.Component {

  renderIndex(data) {
    if(data === null) {
      return;
    }
    const toggle = (children, opened) => {
      if (!children || (children && children.length === 0)) {
        return <span class="prefix"></span>;
      }
      return (
        <span class="prefix" onClick={this._toggleOpened.bind(this)}>
          <ExpandMoreIcon />
        </span>
      );
    }
    const item = i => {
      return (
        <li key={i.name} class={i.opened ? 'opened' : ''}>
          <div>
            {toggle(i.children, i.opened)}
            <span onClick={() => this.props.onItemClick && this.props.onItemClick(i)}>{i.name}</span>
          </div>
          {this.renderIndex(i.children)}
        </li>
      );
    };
    return (
      <ul class="index">
        {data.map(i => item(i))}
      </ul>
   );
  }

  _toggleOpened(e) {
    let node = e.currentTarget;
    while(node === null || (node && node.tagName !== "LI")) {
      node = node.parentNode;
    }
    if(node.classList.contains('opened')) {
      node.classList.remove('opened');
      return;
    }
    node.classList.add('opened');
  }

  render() {
    return (<div class="wrapper">
      {this.renderIndex(this.props.index)}
    </div>);
  }
}
