import React from 'react';
import Page from './Page';
import model from './model';
import {animation} from './animation';
import {BooWrapper, MainCol, SecondCol} from './BooMainWrapper';
import Typography from '@material-ui/core/Typography';
import Menu from '@material-ui/core/Menu';
import IconButton from '@material-ui/core/IconButton';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import BooSticky from './BooSticky';
import BooLink from './BooLink';
import {withStyles} from '@material-ui/core/styles';
import Markdown from 'markdown-to-jsx';
import hljs from 'highlight.js';
import {strings} from './localizer';
import {helper} from './helper';
import {formatter} from './formatter';
import BooIndex from './BlogIndex';
import UserTopBar from './UserTopBar';
import 'highlight.js/styles/monokai.css';


const style = theme => {
  return {
    root: {
      backgroundColor: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      boxShadow: '0px 0px 2px var(--shadow-color)',
      minHeight: '50vh',
      margin: '2px',
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
      '& blockquote': {
        borderLeft: '4px solid var(--blog-hr-color)',
        paddingLeft: '10px',
        overflow: 'auto',
        maxWidth: '100%',
        margin: '10px'
      },
      '& table': {
        display: 'block',
        width: '100%',
        overflow: 'auto'
      },
      '& td': {
        borderBottom: '1px solid var(--blog-hr-color)',
      },
      '& a:hover': {
        textDecoration: 'underline'
      },
      '& div code': {
        color: 'rgb(189, 102, 8)',
        padding: '0px 5px'
      },
      '& p code': {
        color: 'rgb(189, 102, 8)',
        padding: '0px 5px'
      },
      '& li code': {
        color: 'rgb(189, 102, 8)',
        padding: '0px 5px'
      },
      '& ol code': {
        color: 'rgb(189, 102, 8)',
        padding: '0px 5px'
      },
      '& table code': {
        color: 'rgb(189, 102, 8)',
        padding: '0px 5px'
      },
      '& blockquote code': {
        color: 'rgb(189, 102, 8)',
        padding: '0px 5px'
      },
      '& pre code': {
        overflow: 'auto',
        maxWidth: '100%',
        borderRadius: '3px',
        padding: '5px'
      },
      '& img': {
        maxWidth: '100%'
      }
    },
    tag: {
      marginRight: '10px'
    },
    sticky: {
      padding: '0px 8px',
      minHeight: '64px',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      zIndex: 100,
      background: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      '& h1': {
        padding: '0px',
        margin: '0px',
        transition: '.2s all ease-in-out'
      }
    },
    index: {
      maxHeight: 'calc(100vh - 100px)',
      paddingTop: '20px',
      overflow: 'auto'
    },
    stickyRaised: {
      padding: '0px 8px',
      zIndex: 100,
      minHeight: '64px',
      fontSize: '.8em',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      background: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      boxShadow: '0px 0px 5px rgba(0, 0, 0, .3)',
      '& h1': {
        whiteSpace: 'nowrap',
        textOverflow: 'ellipsis',
        overflow: 'hidden',
        padding: '0px',
        margin: '0px',
        transition: '.2s all ease-in-out'
      }
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
      },
      index: null,
      user: {
        name: 'USER NAME',
        avatar: '',
        bio: 'has no bio info',
        contacts: []
      }
    };
    this.content = React.createRef();
    setTimeout(() => {
      this.setState({stickyClass: classes.sticky});
    }, 1);
  }

  enter() {
    const ctx = window.boo.location.context;
    if (this.urlid === ctx.path_params.url_id) {
      return super.enter();
    }
    this.urlid = ctx.path_params.url_id;
    if (!this.inited) {
      this.inited = true;
      window.addEventListener('lang-changed', e => {
        model.userInfo().then(info => {
          this.setState({user:info});
        });
        this.urlid = null;
        if (this.show && ctx.path_params.url_id) {
          this.enter();
        }
      });
      model.userInfo().then(info => {
        this.setState({user: info});
      });
    }
    return model.blog().then(blog => {
      this.setState({blog: blog});
      return new Promise(r => {
        setTimeout(() => {
          helper.updateTitle(this.state.blog.title + ' - ' + this.state.user.name + ' - ' + 'iiiboo');
          helper.updateDescription(this.state.blog.overview);
          helper.updateKeywords(this.state.blog.tags);
          const node = this.content.current;
          if (node) {
            node.querySelectorAll('pre code').forEach(e => {
              hljs.highlightBlock(e);
            });
            const index = helper.index_of_blog(node);
            this.setState({index: index});
          }
          const mj = window.MathJax || false;
          try {
            mj && mj.typeset && mj.typeset();
          } catch(e) {
            console.error(e);
          }
          r();
        }, 10);
      });
    }, (status, msg) => {
      return new Promise(r => r({state: status, msg: msg}));
    }).then(state => {
      return super.enter(state);
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

  handleMenu(e) {
    this.setState({anchorEl: e.currentTarget});
  }

  handleClose(e) {
    this.setState({anchorEl: null});
  }

  indexClicked(item) {
    window.boo.location.go(item.url);
  }

  renderMenu() {
    return (
      <div>
        <IconButton
          aria-label="select a theme to apply"
          aria-controls="menu-appbar"
          aria-haspopup="true"
          onClick={this.handleMenu.bind(this)}
          color="inherit"
        >
          <ExpandMoreIcon />
        </IconButton>
        <Menu
          id="menu-appbar"
          anchorEl={this.state.anchorEl}
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          keepMounted
          transformOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          open={Boolean(this.state.anchorEl)}
          onClose={this.handleClose.bind(this)}
        >
          <div style={{padding: '10px'}}>

            <BooIndex index={this.state.index} onItemClick={this.indexClicked.bind(this)} />
          </div>

        </Menu>
      </div>
    );
  }

  render() {
    const { classes } = this.props;
    return (
      <div style={{padding: '0px 0px 100px 0px'}}>
        <BooWrapper>
          <MainCol>
            <div ref={this.content} className={classes.root}>
              <div style={{padding: '10px 0px', margin: '0px -20px'}}>
                <BooSticky top={0} onRaised={this.onStickyRaised.bind(this)}>
                  <div className={this.state.stickyClass}>
                    <h1>{this.state.blog.title}</h1>
                    {this.renderMenu()}
                  </div>
                </BooSticky>
              </div>
              <Typography component="div" variant="body2" className={classes.row}>
                <label className={classes.label}>{strings.tag}: </label>
                 {this.state.blog.tags.map(t => {
                  return (
                    <BooLink href={'/tags/' + t} className={classes.tag} key={t}>#{t}</BooLink>
                  );
                })}
              </Typography>
              <Typography component="div" variant="body2"  className={classes.row}>
                <label className={classes.label}>{strings.cate}: </label>
                <BooLink href={'/cates/' + this.state.blog.category}>
                  <span>{this.state.blog.category}</span>
                </BooLink>
              </Typography>
              <Typography component="div" variant="body2" className={classes.row}>
                {strings.formatString(strings.edited, {
                  time: formatter.format_time(this.state.blog.updated_at),
                  author: <BooLink href="/">{this.state.user.name}</BooLink>
                })}
              </Typography>
              <hr className={classes.seper} />
              <div ref={this.pc}>
                <Markdown className={classes.content}>
                  {this.state.blog.content}
                </Markdown>
              </div>
            </div>
          </MainCol>
          <SecondCol>
            <div style={{marginTop: '20vh'}}>
              <UserTopBar user={this.state.user}></UserTopBar>
              <BooSticky top={20}>
                <div className={classes.index}>
                    <BooIndex index={this.state.index} onItemClick={this.indexClicked.bind(this)} />
                </div>
              </BooSticky>
            </div>
          </SecondCol>
        </BooWrapper>
      </div>
    )
  }
};

export default withStyles(style)(Blog);
