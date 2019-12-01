import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import LinearProgress from '@material-ui/core/LinearProgress';
import List from './List';
import model from './model';
import Blog from './Blog';
import 'boo-route';

const styles = (theme => {
  return {
    header: {
      height: '40vh',
      backgroundColor: 'black',
      backgroundPosition: 'center center',
      backgroundSize: 'cover',
      backgroundRepeat: 'no-repeat',
      borderBottom: '1px solid #f0f0f0',
      width: '100vw'
    },
    loading: {
      position: 'fixed',
      zIndex: 100,
      left: '0px',
      top: '0px',
      width: '100vw'
    },
    main: {
      position: 'relative',
      top: '-40vh',
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
    app: {
      backgroundColor: '#f0f0f0'
    }
  }
});

const ColorLinearProgress = withStyles({
  colorPrimary: {
    backgroundColor: '#b2dfdb',
  },
  barColorPrimary: {
    backgroundColor: '#00695c',
  },
})(LinearProgress);

class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      pageName: '',
      loading: false
    };
    this.pages = {
      blogs: React.createRef(),
      blog: React.createRef()
    };
    this.header = React.createRef();
  }

  componentDidMount() {
    window.boo.location(ctx => {
      window.boo.route(ctx.pathname, this.routeRules(), (pageName, tail) => {
        this.setState({loading: true});
        this.setPage(pageName).then(() => {
          this.setState({loading: false});
        });
      });
    });
    if ('scrollRestoration' in window.history) {
      window.history.scrollRestoration = 'manual';
    }
    model.userInfo().then(info => {
      this.header.current.style.backgroundImage = 'url('+info.bg+')';
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
      if (this.pageName) {
        const old = this.pages[this.pageName].current;
        return old.exitAnimation().then(old.leave());
      }
      return new Promise(r => r());
    }
    return this.pages[pageName].current.enter(this.pageName).then(() => {
      if (this.pageName === pageName) {
        return new Promise(r => r());
      }
      return leave().then(goto()).then(this.pages[pageName].current.scrollTo()).then(() => {
        this.pageName = pageName;
        return this.pages[this.pageName].current.entryAnimation();
      });
    });
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.app}>
        {this.state.loading && (<ColorLinearProgress className={classes.loading} />)}
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
