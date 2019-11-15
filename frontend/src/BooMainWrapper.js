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
    [theme.breakpoints.up('md')]: {
      width: 'calc(70% - 40px)',
    }
  }
}));

export function MainCol(props) {
  const classes = mainColStyles();
  return <div className={classes.mainCol}>{props.children}</div>
}

const secondColStyles = makeStyles(theme => ({
  secondCol: {
    display: 'none',
    backgroundColor: 'inherit',
    color: 'inherit',
    [theme.breakpoints.up('md')]: {
      display: 'block',
      width: 'calc(30%)'
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
