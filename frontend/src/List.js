import React from 'react'
import {withStyles} from '@material-ui/core/styles';
import UserTopBar from './UserTopBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Grid from '@material-ui/core/Grid';
import {BooWrapper, MainCol, SecondCol } from './BooMainWrapper';
import BlogItem from './BooBlogItem';

const styles = (theme => {
  return {
    mainwrapper: {
      marginTop: '20px'
    }
  };
});

class List extends React.Component {

  constructor(props) {
    super(props);
    window.addEventListener('list-ready', e => {
      this.setState({blogs: e.detail});
    });
    this.state = {value: 0, blogs: []};
  }

  render() {
    const { classes } = this.props;
    return (
      <div>
        <UserTopBar>
          <Tabs
            value={this.state.value}
            indicatorColor="primary"
            textColor="primary"
            aria-label="disabled tabs example"
            variant="scrollable"
            style={{maxWidth: '100%'}}
            scrollButtons="auto" >

            <Tab style={{padding: '6px'}} label="Javascript" />
            <Tab style={{padding: '6px'}} label="PHP" />
            <Tab style={{padding: '6px'}} label="生活琐事" />
            <Tab style={{padding: '6px'}} label="多线程" />
            <Tab style={{padding: '6px'}} label="服务端设计" />
            <Tab style={{padding: '6px'}} label="Active" />
            <Tab style={{padding: '6px'}} label="Active" />
          </Tabs>

        </UserTopBar>
        <div className={classes.mainwrapper}>
          <BooWrapper>
            <MainCol>
              <Grid container spacing={2}>
                {this.state.blogs.map(b => {
                  return (
                    <Grid key={b.url_id} xs={12} item>
                      <BlogItem blog={b}></BlogItem>
                    </Grid>
                  )
                })}
              </Grid>
            </MainCol>
            <SecondCol>
              hello world
            </SecondCol>
          </BooWrapper>
        </div>
      </div>
    );
  }
}

export default withStyles(styles)(List);
