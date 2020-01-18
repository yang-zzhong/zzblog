
import React from 'react';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import {withStyles} from '@material-ui/core/styles';

const style = {
  wrapper: {
    borderLeft: '4px solid var(--card-bg-color)',
    '&>ul': {
      paddingLeft: '5px'
    },
    '& li>ul': {
      display: 'none'
    },
    '& li[opened=true]>ul': {
      display: 'block',
    },
    '& li>ul:hover': {
      cursor: 'pointer',
    },
    '& li>div': {
      display: 'flex',
      alignItems: 'center'
    },
    '& [prefix=true]': {
      display: 'inline-block'
    },
    '& [prefix=true]:hover': {
      cursor: 'pointer'
    },
    '& li>ul:hover': {
      cursor: 'pointer'
    },
    '& li>ul': {
      display: 'none'
    },
    '& li[opened=true]>ul': {
      display: 'block'
    },
    '& [prefix=true]': {
      display: 'inline-block'
    },
    '& [prefix=true]:hover': {
      cursor: 'pointer'
    },
  },
  index: {
    listStyle: 'none',
    paddingLeft: '20px',
    '& li': {
      margin: '4px 0px',
    },
    '& li>div>[prefix=true]': {
      transform: 'rotate(-90deg)',
      transition: 'all .1s ease-in'
    },
    '&>li[opened=true]>div>[prefix=true]': {
      width: '24px',
      height: '24px',
      transition: 'all .2s ease-in-out',
      transform: 'rotate(0deg)'
    }
  },
};

class BooIndex extends React.Component {

  renderIndex(data) {
    if(data === null) {
      return;
    }
    const classes = this.props.classes;
    const toggle = (children, opened) => {
      if (!children || (children && children.length === 0)) {
        return <span prefix="true"></span>;
      }
      return (
        <span prefix="true" onClick={this._toggleOpened.bind(this)}>
          <ExpandMoreIcon />
        </span>
      );
    }
    const item = i => {
      return (
        <li key={i.name} opened={i.opened ? "true" : "false"}>
          <div>
            {toggle(i.children, i.opened)}
            <span onClick={() => this.props.onItemClick && this.props.onItemClick(i)}>{i.name}</span>
          </div>
          {this.renderIndex(i.children)}
        </li>
      );
    };
    return (
      <ul className={classes.index}>
        {data.map(i => item(i))}
      </ul>
   );
  }

  _toggleOpened(e) {
    let node = e.currentTarget;
    while(node === null || (node && node.tagName !== "LI")) {
      node = node.parentNode;
    }
    let opened = node.getAttribute('opened') === 'true';
    node.setAttribute('opened', opened ? "false" : "true");
  }

  render() {
    const {classes} = this.props;
    return (<div className={classes.wrapper}>
      {this.renderIndex(this.props.index)}
    </div>);
  }
}

export default withStyles(style)(BooIndex);
