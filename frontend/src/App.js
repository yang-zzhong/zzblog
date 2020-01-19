import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import LinearProgress from '@material-ui/core/LinearProgress';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import Button from '@material-ui/core/Button';
import ColorLensIcon from '@material-ui/icons/ColorLens';
import ArrowBackIcon from '@material-ui/icons/ArrowBack';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import {localizer} from './localizer';
import {theme} from './theme';
import TranslateIcon from '@material-ui/icons/Translate';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import List from './List';
import Blog from './Blog';
import 'boo-route';

const styles = (theme => {
  return {
    header: {
      height: '40vh',
      backgroundColor: 'var(--head-bg-color)',
      backgroundImage: 'var(--head-bg-img)',
      backgroundPosition: 'center center',
      backgroundSize: 'cover',
      backgroundRepeat: 'no-repeat'
    },
    loading: {
      position: 'fixed',
      zIndex: 10000,
      left: '0px',
      top: '0px',
      width: '100vw'
    },
    main: {
      position: 'relative',
      top: '-45vh',
      [theme.breakpoints.up('xs')]: {
        left: '2%',
        width: '96%'
      },
      [theme.breakpoints.up('sm')]: {
        left: 'calc(50vw - 250px)',
        width: '500px'
      },
      [theme.breakpoints.up('md')]: {
        left: 'calc(50vw - 400px)',
        width: '800px'
      },
      [theme.breakpoints.up('lg')]: {
        left: 'calc(50vw - 500px)',
        width: '1000px'
      }
    },
    mainwrapper: {
      marginTop: '30vh',
      minHeight: '70vh',
      paddingBottom: '50px',
      position: 'relative',
      width: '100%'
    },
    page: {
      position: 'absolute',
      width: '100%',
      display: 'none',
      '&[selected=true]': {
        display: 'block'
      }
    },
    appbar: {
      position: 'absolute',
      backgroundColor: 'transparent',
      color: "var(--head-fg-color)",
      boxShadow: 'none'
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
    app: {
      position: 'relative',
      backgroundColor: 'var(--bg-color)',
      olor: 'var(--fg-color)'
    }
  }
});

const ColorLinearProgress = withStyles({
  colorPrimary: {
    backgroundColor: 'var(--loading-bg-color)',
  },
  barColorPrimary: {
    backgroundColor: 'var(--loading-fg-color)',
  },
})(LinearProgress);

class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      pageName: '',
      loading: false,
      langLabel: '',
      theme: [],
      langs: []
    };
    window.addEventListener('lang-changed', e => {
      this.setLangLabel(e.detail);
    });
    this.pages = {
      blogs: React.createRef(),
      blog: React.createRef()
    };
    this.header = React.createRef();
    theme.init().then(data => {
      this.setState({theme: data});
      let t = theme.guess();
      if (t) {
        theme.use(t);
      }
    });
  }

  handleLang(e) {
    this.setState({anchorLang: e.currentTarget});
  }

  langClose(lang) {
    this.setState({anchorLang: null});
    if (lang) {
      localizer.use(lang);
    }
  }

  handleMenu(e) {
    this.setState({anchorEl: e.currentTarget});
  };

  handleClose(e, key) {
    theme.use(key);
    this.setState({anchorEl: null});
  };

  componentDidMount() {
    window.boo.location(ctx => {
      window.boo.route(ctx.pathname, this.routeRules(), this.route.bind(this));
    });
    if ('scrollRestoration' in window.history) {
      window.history.scrollRestoration = 'manual';
    }
  }

  setLangLabel(lang) {
    const setLangState = lang => {
      for (let i = 0; i < this.state.langs.length; ++i) {
        if (this.state.langs[i].name === lang) {
          this.setState({langLabel: this.state.langs[i].label});
          return true;
        }
      }
      return false;
    };
    if (setLangState(lang)) {
      return;
    }
    let la = lang.split('-');
    console.log(la);
    setLangState(la[0]);
  }

  route(pageName, tail) {
    localizer.ready().then(langs => {
      this.setState({langs: langs});
      if (this.lang === undefined) {
        this.lang = localizer.guess();
        console.log(this.lang);
        localizer.use(this.lang);
      }
      return new Promise(r => r());
    }).then(() => {
      let ps = [new Promise(r => {
        setTimeout(() => {
          r();
        }, 0);
      })];
      this.setState({loading: true});
      ps.push(this.setPage(pageName));
      Promise.all(ps).then(() => {
        this.setState({loading: false});
      });
    });
  }

  routeRules() {
    return [
      { rule: "/cates/:catename/tags/:tagname", page: "blogs" },
      { rule: "/cates/:catename", page: "blogs" },
      { rule: "/tags/:tagname", page: "blogs" },
      { rule: "/:url_id", page: "blog" }
    ];
  }

  setPage(pageName) {
    if (pageName === '__root') {
      pageName = 'blogs';
    }
    const goto = () => {
      let pages = this.refs.wrapper.querySelectorAll('[page=true]');
      pages.forEach(page => {
        if (page.getAttribute('id') !== pageName) {
          page.removeAttribute('selected');
        } else {
          page.setAttribute('selected', 'true');
        }
      });
      return new Promise(r => r());
    };
    const leave = () => {
      if (this.state.pageName) {
        const old = this.pages[this.state.pageName].current;
        return old.exitAnimation().then(old.leave());
      }
      return new Promise(r => r());
    }
    return this.pages[pageName].current.enter(this.state.pageName).then(() => {
      if (this.state.pageName === pageName) {
        return this.pages[pageName].current.scrollTo();
      }
      return leave().then(goto()).then(this.pages[pageName].current.scrollTo()).then(() => {
        this.setState({pageName: pageName});
        return this.pages[this.state.pageName].current.entryAnimation();
      });
    });
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.app}>
        {this.state.loading && (<ColorLinearProgress className={classes.loading} />)}
        <AppBar position="static" className={classes.appbar}>
          <Toolbar>
            {this.state.pageName === "blog" && (
              <IconButton onClick={() => window.history.go(-1)} style={{color: "var(--head-fg-color)"}}>
                <ArrowBackIcon />
              </IconButton>
            )}
            <Typography variant="h6" className={classes.title}></Typography>
            <div>
              <Button
                aria-label="select a theme to apply"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={this.handleLang.bind(this)}
                color="inherit"
              >
                <TranslateIcon />
                <span style={{display: 'inline-block', margin: '0px 5px'}}>{this.state.langLabel}</span>
                <ExpandMoreIcon />
              </Button>
              <Menu
                id="menu-appbar"
                anchorEl={this.state.anchorLang}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'center',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'center',
                }}
                open={Boolean(this.state.anchorLang)}
                onClose={() => this.langClose()}
              >
                {this.state.langs.map(l => {
                  return <MenuItem key={l.name} onClick={() => this.langClose(l.name)}>{l.label}</MenuItem>
                })}
              </Menu>
            </div>
            <div>
              <IconButton
                aria-label="select a theme to apply"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={this.handleMenu.bind(this)}
                color="inherit"
              >
                <ColorLensIcon />
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

                {this.state.theme.map(t => {
                  return <MenuItem key={t.name} onClick={(e) => this.handleClose(e, t.name)}>{t.name}</MenuItem>
                })}
              </Menu>
            </div>
          </Toolbar>
        </AppBar>
        <header ref={this.header} className={classes.header}></header>
        <div className={classes.main}>
          <main className={classes.mainwrapper} ref="wrapper">
            <div className={classes.page} page="true" id="blogs">
              <List ref={this.pages.blogs} />
            </div>
            <div className={classes.page} page="true" id="blog">
              <Blog ref={this.pages.blog} />
            </div>
          </main>
        </div>
      </div>
    );
  }
};

export default withStyles(styles)(App);
