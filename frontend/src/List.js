import React from 'react'
import Page from './Page';
import {animation} from './animation';
import {withStyles} from '@material-ui/core/styles';
import  CircularProgress from '@material-ui/core/CircularProgress';
import UserTopBar from './UserTopBar';
import SweetScroll from 'sweet-scroll';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Grid from '@material-ui/core/Grid';
import {BooWrapper, MainCol, SecondCol } from './BooMainWrapper';
import Button from '@material-ui/core/Button';
import BlogItem from './BooBlogItem';
import model from './model';
import BooLink from './BooLink';

const styles = (theme => {
  return {
    mainwrapper: {
      marginTop: '20px'
    },
    nomorewrapper: {
      padding: '10px',
      textAlign: 'center'
    },
    tagrapper: {
      backgroundColor: 'white',
      boxShadow: '0px 0px 2px rgba(0, 0, 0, .3)',
      borderRadius: '3px',
      padding: '10px',
      display: 'flex',
      justifyContent: 'flex-start',
      alignItems: 'flex-start',
      flexWrap: 'wrap'
    },
    tag: {
      padding: '5px',
      display: 'block',
      whiteSpace: 'nowrap'
    },
    selectedtag: {
      padding: '5px',
      display: 'block',
      whiteSpace: 'nowrap',
      color: 'red'
    }
  };
});

class List extends Page {

  constructor(props) {
    super(props);
    model.cates().then(cates => {
      cates.unshift('全部');
      this.setState({cates: cates});
      this.updateSelected();
    });
    model.tags().then(tags => {
      this.setState({tags: tags});
    });
    this.state = {
      value: 0,
      loading: false,
      blogs: [],
      tags: [],
      cates: ['全部'],
      noMore: false
    };
    this.anis = {
      user: React.createRef(),
    };
  }

  enter(old) {
    let params = model.queryBlogParams();
    if (!params) {
      return super.enter();
    }
    this.cate = params.cate ? params.cate : '全部';
    this.tag = decodeURIComponent(params.tag);
    this.page = params.page;
    this.updateSelected();
    this.setState({loading: true, blogs: []});
    return model.queryBlogs(params).then(blogs => {
      this.setState({
        loading: false,
        noMore: blogs.length === 0,
        blogs: blogs || []
      });
      if (old === 'blogs') {
        const scroll = new SweetScroll();
        scroll.toTop();
        const t = this.blogsInAni();
        if (t) {
          animation.play(t);
        }
      }
      return super.enter();
    });
  }

  leave() {
    this.anis.user.current.disableSticky(true);
    return super.leave();
  }

  aniBlogs(begin) {
    const bs = [];
    const blogs = this.pc.current.querySelectorAll('[item=true]');
    for(let i = begin; i < blogs.length; ++i) {
      if (this.isVisible(blogs[i])) {
        bs.push(blogs[i]);
      }
    }
    return bs;
  }

  blogsInAni(begin) {
    const bs = this.aniBlogs(begin || 0);
    if (bs.length === 0) {
      return false;
    }
    const topins = animation.bottom_in(bs);
    topins.delay = 100;
    return topins;
  }

  blogsOutAni() {
    const bs = this.aniBlogs();
    if (bs.length > 0) {
      const topins = animation.top_out(bs);
      topins.delay = 100;
      return topins;
    }
    return false;
  }

  entryAnimation() {
    const node = this.anis.user.current.root();
    const anis = [ animation.scale_in([node]) ];
    const t = this.blogsInAni();
    if (t) {
      anis.push(t);
    }
    setTimeout(() => {
      this.anis.user.current.disableSticky(false);
      this.anis.user.current.updateSticky();
    }, 500);
    return this.playAnimation(anis);
  }

  // exitAnimation() {
  //   const node = this.anis.user.current;
  //   const anis = [ animation.scale_out([node]) ];
  //   const t = this.blogsOutAni();
  //   if (t) {
  //     anis.push(t);
  //   }
  //   return this.playAnimation(anis);
  // }

  updateSelected() {
    for (let i = 0; i < this.state.cates.length; ++i) {
      if (this.state.cates[i] === this.cate) {
        this.setState({value: i});
        break;
      }
    }
  }

  handleChange(e, newValue) {
    let url = "";
    if (newValue !== 0 && newValue < this.state.cates.length) {
      url = "/cates/" + this.state.cates[newValue];
    }
    window.boo.location.go(url);
  }

  render() {
    const { classes } = this.props;
    return (
      <div ref={this.pc}>
        <UserTopBar ref={this.anis.user}>
          <Tabs
            value={this.state.value}
            indicatorColor="primary"
            textColor="primary"
            aria-label="disabled tabs example"
            onChange={this.handleChange.bind(this)}
            variant="scrollable"
            style={{maxWidth: '100%'}}
            scrollButtons="auto" >

            {this.state.cates.map(cate => {
              return <Tab key={cate} label={cate} />;
            })}

          </Tabs>

        </UserTopBar>
        <div className={classes.mainwrapper}>
          <BooWrapper>
            <MainCol>
              <Grid container spacing={2}>
                {this.state.blogs.map(b => {
                  return (
                    <Grid key={b.url_id + Math.random()} xs={12} item>
                      <div item="true">
                        <BlogItem blog={b}></BlogItem>
                      </div>
                    </Grid>
                  )
                })}
              </Grid>
              <div className={classes.nomorewrapper}>
                {this.moreButton()}
              </div>
            </MainCol>
            <SecondCol>
              <div className={classes.tagrapper}>
                {this.state.tags.map(b => {
                  return (
                    <BooLink key={b} className={b === this.tag ? classes.selectedtag : classes.tag} href={"/tags/" + b}>{'#' + b}</BooLink>
                  );
                })}
              </div>
            </SecondCol>
          </BooWrapper>
        </div>
      </div>
    );
  }

  loadMore() {
    this.setState({loading: true});
    model.nextBlogs().then(blogs => {
      const b = this.state.blogs.length;
      blogs = this.state.blogs.concat(blogs || []);
      this.setState({
        noMore: blogs.length === 0,
        blogs: blogs,
        loading: false
      });
      const items = this.pc.current.querySelectorAll('[item=true]');
      const scroll = new SweetScroll();
      if (items.length > b) {
        scroll.toElement(items[b]);
      }
      for(let i = b; i < items.length; ++i) {
        items[i].style.visibility = 'hidden';
      }
      setTimeout(() => {
        const items = this.pc.current.querySelectorAll('[item=true]');
        for(let i = b; i < items.length; ++i) {
          items[i].style.visibility = 'visible';
        }
        const t = this.blogsInAni(b);
        if (t) {
          return animation.play(t).then(super.enter());
        }
      }, 300);
    });
  }

  moreButton() {
    if (this.state.loading) {
      const ColorCircularProgress = withStyles({
        root: {
          color: '#00695c',
        },
      })(CircularProgress);
      return <ColorCircularProgress />;
    }
    if (this.state.noMore) {
      return <span>没有跟多内容了</span>;
    }
    return <Button onClick={this.loadMore.bind(this)}>加载更多</Button>;
  }
}

export default withStyles(styles)(List);
