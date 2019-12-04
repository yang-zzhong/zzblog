import React from 'react';
import Page from './Page';
import model from './model';
import {animation} from './animation';
import {BooWrapper, MainCol} from './BooMainWrapper';
import BooSticky from './BooSticky';
import BooLink from './BooLink';
import {withStyles} from '@material-ui/core/styles';
import Markdown from 'markdown-to-jsx';
import hljs from 'highlight.js';
import {strings} from './localizer';
import {formatter} from './formatter';
import 'highlight.js/styles/github.css';


const style = theme => {
  return {
    root: {
      backgroundColor: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      boxShadow: '1px 1px 20px var(--shadow-color)',
      minHeight: '50vh',
      padding: '20px'
    },
    row: {
      backgroundColor: 'var(--card-bg-color)',
      marginBottom: '5px',
      marginTop: '5px'
    },
    seper: {
      border: '1px solid var(--blog-hr-color)',
    },
    label: {
      marginRight: '10px',
    },
    content: {
      backgroundColor: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      lineHeight: '1.5',
      '& h1': {
        color: 'var(--blog-h-color)',
      },
      '& h2': {
        color: 'var(--blog-h-color)',
      },
      '& h3': {
        color: 'var(--blog-h-color)',
      },
      '& h4': {
        color: 'var(--blog-h-color)',
      },
      '& h5': {
        color: 'var(--blog-h-color)',
      },
      '& h6': {
        color: 'var(--blog-h-color)',
      },
      '& hr': {
        border: '1px solid var(--blog-hr-color)'
      },
      '& a': {
        color: 'var(--blog-a-color)',
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
    },
    sticky: {
      margin: '0px -20px',
      padding: '20px',
      width: '100%',
      zIndex: 100,
      background: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)'
    },
    stickyRaised: {
      margin: '0px -20px',
      padding: '20px',
      width: 'clac(100% + 20px)',
      zIndex: 100,
      background: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      boxShadow: '0px 0px 5px rgba(0, 0, 0, .3)',
      whiteSpace: 'nowrap',
      textOverflow: 'ellipsis',
      overflow: 'hidden'
    },
  };
};

class Blog extends Page {

  constructor(props) {
    super(props)
    const {classes} = this.props;
    this.state = {
      blog: {
        stickyClass: classes.sticky,
        title: "title",
        content: "content",
        category: "category",
        tags: []
      }
    };
    window.addEventListener('lang-changed', e => {
      if (this.show) {
        this.urlid = null;
        this.enter();
      }
    });
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

  onStickyRaised(raised) {
    const {classes} = this.props;
		if (raised) {
      this.setState({stickyClass: classes.stickyRaised});
		} else {
      this.setState({stickyClass: classes.sticky});
		}
  }

  render() {
    const { classes } = this.props;
    return (
      <div ref={this.content} className={classes.root}>
        <BooWrapper>
          <MainCol>
            <BooSticky top={0} onRaised={this.onStickyRaised.bind(this)}>
              <h1 className={this.state.stickyClass}>{this.state.blog.title}</h1>
            </BooSticky>
            <div className={classes.row}><label className={classes.label}>{strings.tag}: </label>{this.state.blog.tags.map(t => {
            return (
              <BooLink href={'/tags/' + t} className={classes.tag} key={t}>#{t}</BooLink>
            );
            })}</div>
            <div className={classes.row}>
              <label className={classes.label}>{strings.cate}: </label>
              <BooLink href={'/cates/' + this.state.blog.category}>
                <span>{this.state.blog.category}</span>
              </BooLink>
            </div>
            <div className={classes.row}>
              {formatter.format_time(this.state.blog.updated_at)} {strings.edited}
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
