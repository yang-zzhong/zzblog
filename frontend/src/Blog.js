import React from 'react';
import {BooWrapper, MainCol} from './BooMainWrapper';
import {withStyles} from '@material-ui/core/styles';

const style = theme => {
  return {
    root: {
      backgroundColor: 'white',
      minHeight: '50vh',
      padding: '20px'
    }
  };
};

class Blog extends React.Component {

  constructor(props) {
    super(props)
    this.state = {};
  }

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.root}>
        <BooWrapper>
          <MainCol>
            <h1>这点是标题</h1>
          </MainCol>
        </BooWrapper>
      </div>
    )
  }
};

export default withStyles(style)(Blog);
