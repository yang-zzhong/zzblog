import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import {formatter} from './formatter';
import {strings} from './localizer';

const useStyles = makeStyles(theme =>({
  card: {
    borderRadius: '0px',
    padding: '5px',
    backgroundColor: 'var(--card-bg-color)',
    color: 'var(--card-fg-color)',
    boxShadow: '1px 1px 3px var(--shadow-color)',
    '&:hover': {
      cursor: 'pointer'
    }
  },
  img: {
    maxHeight: '100%',
    width: 'auto',
    borderRadius: '2px',
    maxWidth: '50%',
    [theme.breakpoints.down('sm')]: {
      width: 'calc(100% - 25px)',
      maxWidth: '100%'
    }
  }
}));

export default function BlogItem(props) {
  const classes = useStyles();

  return (
    <Card className={classes.card}>
      <CardContent className={classes.area} onClick={() => window.boo.location.go("/" + props.blog.url_id)}>
        <Typography gutterBottom variant="h5" component="h2">
          <span style={{color: 'var(--blog-h-color)'}}>{props.blog.title}</span>
        </Typography>
        <Typography variant="body2" component="p">
          {Boolean(props.blog.image) && (
            <img className={classes.img} alt="img" align="left" hspace="10" vspace="3" src={props.blog.image} />
          )}
          <span style={{color: 'var(--card-fg-color)'}}>{props.blog.overview}</span>
        </Typography>
      </CardContent>
      <CardActions>
        <Typography gutterBottom variant="overline" component="span">
          {formatter.format_time(props.blog.created_at)} {strings.edited}
        </Typography>
      </CardActions>
    </Card>
  );
}
