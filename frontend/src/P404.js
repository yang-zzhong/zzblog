import React from 'react';
import Page from './Page';
import {BooWrapper, MainCol} from './BooMainWrapper';
import {withStyles} from '@material-ui/core/styles';
import {strings} from './localizer';
import BooLink from './BooLink';


const style = theme => {
  return {
    root: {
      backgroundColor: 'var(--card-bg-color)',
      color: 'var(--card-fg-color)',
      boxShadow: '0px 0px 2px var(--shadow-color)',
      minHeight: '50vh',
      margin: '2px',
      padding: '20px'
    }
  };
};

class P404 extends Page {

  render() {
    const { classes } = this.props;
    return (
      <div style={{padding: '0px 0px 100px 0px'}}>
        <BooWrapper>
          <MainCol>
            <div className={classes.root}>
              <h2>{strings.blognotfound}</h2>
              <BooLink href='/'>
                <span>{strings.gotohome}</span>
              </BooLink>
            </div>
          </MainCol>
        </BooWrapper>
      </div>
    )
  }
};

export default withStyles(style)(P404);
