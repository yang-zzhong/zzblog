import React from 'react'
import Page from './Page';
import {animation} from './animation';
import {withStyles} from '@material-ui/core/styles';
import {helper} from './helper';
import  CircularProgress from '@material-ui/core/CircularProgress';
import BooSticky from './BooSticky';
import {strings} from './localizer';
import Chip from '@material-ui/core/Chip';
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
      marginTop: '30px'
    },
    nomorewrapper: {
      padding: '10px',
      textAlign: 'center'
    },
    tagrapper: {
      backgroundColor: 'var(--card-bg-color)',
      boxShadow: '1px 1px 3px var(--shadow-color)',
      borderRadius: '0px',
      padding: '10px',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'flex-start',
      maxHeight: 'calc(100vh - 120px)',
      overflow: 'auto',
      flexWrap: 'wrap'
    },
    tag: {
      padding: '2px 5px',
      display: 'block',
      whiteSpace: 'nowrap'
    },
    userTopBar: {
      marginLeft: '200px',
      [theme.breakpoints.up('xl')]: {
        marginLeft: '200px'
      }
    },
    selected: {
      color: 'red'
    }
  };
});

class List extends Page {

  constructor(props) {
    super(props);
    this.state = {
      value: 0,
      loading: false,
      blogs: [],
      tags: [],
      user: {
        name: 'USER NAME',
        avatar: '',
        bio: 'has no bio info',
        contacts: []
      },
      cates: [],
      noMore: false
    };
    this.anis = {
      user: React.createRef(),
      grid: React.createRef()
    };
  }

  tagAndCate() {
    model.tags().then(tags => {
      this.setState({tags: tags});
      return new Promise(r => r());
    });
    return model.cates().then(cates => {
      this.setState({cates: cates});
      this.updateSelected();
      return new Promise(r => r());
    });
  }

  enter(old) {
    let params = model.queryBlogParams();
    if (!params) {
      return super.enter();
    }
    if (!this.inited) {
      this.inited = true;
      window.addEventListener('lang-changed',  e => {
        model.clearOldBlogQuery();
        this.tagAndCate();
        model.userInfo().then(info => {
          this.setState({user: info});
        });
        if (this.show) {
          window.boo.location.replace('/');
        }
      });
    }
    if (this.tagAndCateReady) {
      return this.loadBlogs(params, old).then(r => {
        return super.enter();
      });
    }
    this.tagAndCateReady = true;
    return Promise.all([
      this.tagAndCate(),
      model.userInfo().then(info => {
        this.setState({user: info});
        return new Promise(r => r());
      }),
      this.loadBlogs(params, old)
    ]).then(() => {
      const info = this.state.user;
      helper.updateTitle(info.name + ' - Home Page');
      let desc = [info.name + '\'s home page. ', 'bio: ' + info.bio];
      for(let i =0; i < info.contacts.length; ++i) {
        desc.push(info.contacts[i].label + ': ' + info.contacts[i].value);
      }
      helper.updateDescription(desc.join('\n'));
      let keywords = ['iiiboo', info.name];
      for(let i = 0; i < this.state.cates.length; ++i) {
        keywords.push(this.state.cates[i]);
      }
      helper.updateKeywords(keywords);
      return super.enter(old);
    });
  }

  loadBlogs(params, old) {
    this.cate = decodeURIComponent(params.cate ? params.cate : 'all');
    this.tag = decodeURIComponent(params.tag);
    this.page = params.page;
    this.updateSelected();
    this.setState({loading: true, blogs: []});
    return model.queryBlogs(params).then(blogs => {
      this.afterLoadedBlogs(blogs);
      if (old === 'blogs') {
        this.anis.grid.current.style.visibility = 'hidden';
        setTimeout(() => {
          this.anis.grid.current.style.visibility = 'visible';
          const scroll = new SweetScroll();
          scroll.toTop();
          const t = this.blogsInAni();
          if (t) {
            animation.play(t);
          }
        }, 20);
      }
      return new Promise(r => r());
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
      bs.push(blogs[i]);
    }
    return bs;
  }

  blogsInAni(begin) {
    const bs = this.aniBlogs(begin || 0);
    if (bs.length === 0) {
      return false;
    }
    const topins = animation.bottom_in(bs);
    topins.delay = 30;
    return topins;
  }

  blogsOutAni() {
    const bs = this.aniBlogs();
    if (bs.length > 0) {
      const topins = animation.top_out(bs);
      topins.delay = 30;
      return topins;
    }
    return false;
  }

  aniElems() {
    const node = this.anis.user.current.root();
    const anis = [ animation.scale_in([node]) ];
    const blogs = this.pc.current.querySelectorAll('[item=true]');
    blogs.forEach((item) => {
      anis.push(item);
    });
    return anis;
  }

  hiddenElems(elems, hidden) {
    for(let i in elems) {
      if (elems[i].style) {
        elems[i].style.visibility = hidden ? 'hidden' : 'visible';
      }
    }
  }

  entryAnimation() {
    this.hiddenElems(this.aniElems(), true);
    return new Promise(r => {
      setTimeout(() => {
        const node = this.anis.user.current.root();
        const anis = [ animation.scale_in([node]) ];
        const t = this.blogsInAni();
        if (t) {
          anis.push(t);
        }
        this.anis.user.current.disableSticky(true);
        setTimeout(() => {
          this.anis.user.current.disableSticky(false);
          this.anis.user.current.updateSticky();
        }, 500);
        r(anis);
        this.hiddenElems(this.aniElems(), false);
      }, 10);
    }).then(anis => this.playAnimation(anis));
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
        this.setState({value: i + 1});
        return;
      }
    }
    this.setState({value: 0});
  }

  handleChange(e, newValue) {
    let url = "/";
    if (newValue > 0 && newValue <= this.state.cates.length) {
      url = "/cates/" + this.state.cates[newValue - 1];
    }
    window.boo.location.go(url);
  }

  render() {
    const { classes } = this.props;
    return (
      <div ref={this.pc}>
        <UserTopBar className={classes.userTopBar} user={this.state.user} ref={this.anis.user}>
          <Tabs
            value={this.state.value}
            aria-label="disabled tabs example"
            onChange={this.handleChange.bind(this)}
            variant="scrollable"
            style={{
              maxWidth: '100%',
              background: 'var(--tab-bg-color)',
              color: 'var(--tab-fg-color)',
              indicator: {
                backgroundColor: 'var(--tab-ind-color)',
                color: 'var(--tab-ind-color)'
              }}}
            scrollButtons="auto" >
            <Tab key='all' label={strings.all} />
            {this.state.cates.map(cate => {
              return <Tab key={cate} label={cate} />;
            })}

          </Tabs>

        </UserTopBar>
        <div className={classes.mainwrapper}>
          <BooWrapper>
            <MainCol>
              <div style={{padding: '2px'}}>
                <Grid ref={this.anis.grid} container spacing={2}>
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
              </div>
            </MainCol>
            <SecondCol>
              <BooSticky top={64}>
                <div className={classes.tagrapper}>
                  {this.state.tags.map(b => {
                    return (
                      <BooLink key={b} className={classes.tag} href={"/tags/" + b}>
                        <Chip
                          label={b}
                          color={b === this.tag ? "secondary" : "primary"}
                          variant="outlined"/>
                      </BooLink>
                    );
                  })}
                </div>
              </BooSticky>
            </SecondCol>
          </BooWrapper>
        </div>
      </div>
    );
  }

  loadMore() {
    this.setState({loading: true});
    model.nextBlogs().then(blogs => this.afterLoadedBlogs(blogs, true));
  }

  afterLoadedBlogs(blogs, scroll) {
    const b = this.state.blogs.length;
    this.setState({
      noMore: blogs.length < 10,
      blogs: this.state.blogs.concat(blogs || []),
      loading: false
    });
    if (scroll) {
      const items = this.pc.current.querySelectorAll('[item=true]');
      const scroll = new SweetScroll();
      if (items.length > b) {
        scroll.toElement(items[b]);
      }
      this.hiddenElems(items, true);
      setTimeout(() => {
        const items = this.pc.current.querySelectorAll('[item=true]');
        this.hiddenElems(items, false);
        const t = this.blogsInAni(b);
        if (t) {
          return animation.play(t).then(super.enter());
        }
      }, 300);
    }
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
      return <span>{strings.noMore}</span>;
    }
    return <Button onClick={this.loadMore.bind(this)}><span>{strings.loadMore}</span></Button>;
  }
}

export default withStyles(styles)(List);
