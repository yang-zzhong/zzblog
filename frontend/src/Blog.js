import React from 'react';
import Page from './Page';
import model from './model';
import {animation} from './animation';
import {BooWrapper, MainCol} from './BooMainWrapper';
import BooLink from './BooLink';
import {withStyles} from '@material-ui/core/styles';
import Markdown from 'markdown-to-jsx';
import hljs from 'highlight.js';
import 'highlight.js/styles/github.css';


const style = theme => {
  return {
    root: {
      backgroundColor: 'white',
      boxShadow: '0px 0px 2px rgba(0, 0, 0, .3)',
      minHeight: '50vh',
      padding: '20px'
    },
    row: {
      marginBottom: '5px',
      marginTop: '5px'
    },
    seper: {
      border: '1px solid #f0f0f0',
    },
    label: {
      marginRight: '10px',
    },
    content: {
      color: '#444',
      lineHeight: '1.5',
      '& h1': {
        color: 'black',
      },
      '& h2': {
        color: 'black',
      },
      '& h3': {
        color: 'black',
      },
      '& h4': {
        color: 'black',
      },
      '& h5': {
        color: 'black',
      },
      '& h6': {
        color: 'black',
      },
      '& hr': {
        border: '1px solid #f0f0f0'
      },
      '& a': {
        color: '#1a0dab',
        textDecoration: 'none'
      },
      '& a:hover': {
        textDecoration: 'underline'
      },
      '& code': {
        overflow: 'auto',
        maxWidth: '100%'
      },
      '& img': {
        maxWidth: '100%'
      }
    },
    tag: {
      marginRight: '10px'
    }
  };
};

class Blog extends Page {

  constructor(props) {
    super(props)
    this.state = {
      blog: {
        title: "title",
        content: "content",
        category: "category",
        tags: []
      }
    };
    this.content = React.createRef();
  }

  enter() {
    const ctx = window.boo.location.context;
    if (this.urlid === ctx.path_params.url_id) {
      return super.enter();
    }
    return model.blog().then(blog => {
      this.setState({blog: blog});
      setTimeout(() => {
        const node = this.content.current;
        if (node) {
          node.querySelectorAll('pre code').forEach(e => {
            hljs.highlightBlock(e);
          });
        }
      }, 100);
      return super.enter();
    });
  }

  entryAnimation() {
    const node = this.content.current;
    if (!node) {
      return new Promise(r => r());
    }
    let anis = [animation.fade_in([node])];
    let nodes = [node];
    let elems = node.querySelectorAll('h1, h2, ul, li, ol, h3, h4, h5, h6, div, img, p, pre code');
    for (let i = 0; i < elems.length; ++i) {
      if (this.isVisible(elems[i])) {
        nodes.push(elems[i]);
      }
    }
    if (nodes.length > 0) {
      let ani = animation.top_in(nodes);
      ani.delay = 10;
      anis.push(ani);
    }
    return this.playAnimation(anis);
  }

  exitAnimation() {
    const node = this.content.current;
    return animation.play(animation.top_out([node]));
  }

  render() {
    const { classes } = this.props;
    return (
      <div ref={this.content} className={classes.root}>
        <BooWrapper>
          <MainCol>
            <h1>{this.state.blog.title}</h1>
            <div className={classes.row}><label className={classes.label}>标签: </label>{this.state.blog.tags.map(t => {
            return (
              <BooLink href={'/tags/' + t} className={classes.tag} key={t}>#{t}</BooLink>
            );
            })}</div>
            <div className={classes.row}>
              <label className={classes.label}>分类: </label>
              <BooLink href={'/cates/' + this.state.blog.category}>
                <span>{this.state.blog.category}</span>
              </BooLink>
            </div>
            <div className={classes.row}>
            </div>
            <hr className={classes.seper} />
            <Markdown className={classes.content}>
              {this.state.blog.content}
            </Markdown>
          </MainCol>
        </BooWrapper>
      </div>
    )
  }
};

export default withStyles(style)(Blog);
