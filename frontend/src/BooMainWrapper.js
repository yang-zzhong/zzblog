import React from 'react';
import {makeStyles} from '@material-ui/core/styles';

const wrapperStyles = makeStyles(theme => ({
  content: {
    display: 'flex',
    alignItems: 'flex-start',
    justifyContent: 'space-between',
    width: '100%',
    color: 'inherit',
    background: 'inherit'
  }
}));

export function BooWrapper(props) {
  const classes = wrapperStyles();
  return (
    <div className={classes.content}>
      {props.children}
    </div>
  );
};

const mainColStyles = makeStyles(theme => ({
  mainCol: {
    display: 'block',
    width: '100%',
    color: 'inherit',
    background: 'inherit',
  }
}));

export const MainCol = React.forwardRef((props, ref) => {
  const classes = mainColStyles();
  return <div className={classes.mainCol} ref={ref}>{props.children}</div>
});

const secondColStyles = makeStyles(theme => ({
  secondCol: {
    display: 'none',
    backgroundColor: 'inherit',
    color: 'inherit',
    [theme.breakpoints.up('md')]: {
      display: 'block',
      paddingLeft: '20px',
      width: '400px'
    },
    [theme.breakpoints.up('lg')]: {
      display: 'block',
      paddingLeft: '20px',
      width: '500px'
    },
    [theme.breakpoints.up('xl')]: {
      display: 'block',
      paddingLeft: '20px',
      width: '600px'
    }
  }
}));

export function SecondCol(props) {
  const classes = secondColStyles();
  return <div className={classes.secondCol}>{props.children}</div>
}

export function BooContent(props) {
  return <div style={{
    marginTop: props.marginTop || '100px',
    backgroundColor: 'inherit',
    color: 'inherit'}}>{props.children}</div>
}
