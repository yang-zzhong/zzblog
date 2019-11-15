import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import List from './List';
import Blog from './Blog';
import model from './model';
import 'boo-route';

const styles = (theme => {
  return {
    header: {
      height: '40vh',
      backgroundColor: 'black',
      backgroundImage: 'url(https://i.ytimg.com/vi/g55cDsjzitg/maxresdefault.jpg)',
      backgroundPosition: 'center center',
      backgroundSize: 'cover',
      backgroundRepeat: 'no-repeat',
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

class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      pageName: ''
    };
    this.list = React.createRef();
    this.blog = React.createRef();
  }

  componentDidMount() {
    window.boo.location(ctx => {
      window.boo.route(ctx.pathname, this.routeRules(), (pageName, tail) => {
        this.setPage(pageName);
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
    }
    if (pageName === 'blogs') {
      model.blogs().then(() => goto());
    } else if (pageName === 'blog') {
      model.blog().then(() => goto());
    }
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.app}>
        <header className={classes.header}>
        </header>
        <div className={classes.main}>
          <main className={classes.mainwrapper} ref="wrapper">
            <div className={classes.page} page="true" id="blogs">
              <List ref={this.list} />
            </div>
            <div className={classes.page} page="true" id="blog">
              <Blog ref={this.blog} />
            </div>
          </main>
        </div>
      </div>
    );
  }
};

export default withStyles(styles)(App);
