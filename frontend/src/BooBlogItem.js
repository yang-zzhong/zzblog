import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';
import {formatter} from './formatter';
import {strings} from './localizer';
import model from './model';

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
    maxHeight: '150px',
    width: 'auto',
    visibility: 'hidden',
    [theme.breakpoints.down('sm')]: {
      width: '100%',
      maxWidth: '100%'
    }
  },
  content: {
    display: 'flex',
    justifyContent: 'flex-start',
    alignItems: 'flex-start',
    flexWrap: 'no-wrap',
    [theme.breakpoints.down('sm')]: {
      display: 'block',
      justifyContent: 'none',
      alignItems: 'none',
      flexWrap: 'wrap',
    }
  },
  imgwrapper: {
    marginRight: '10px',
    marginBottom: '10px',
    backgroundPosition: 'center center',
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover',
    borderRadius: '2px',
    height: 'auto',
    display: 'block',
    maxWidth: '40%',
    [theme.breakpoints.down('sm')]: {
      width: '100%',
      marginRight: '0px',
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
        <Typography variant="body2" className={classes.content} component="p">
          <span className={classes.imgwrapper} style={{backgroundImage:"url(" + model.imageUrl(props.blog.image, 300) + ")"}}>
            {Boolean(props.blog.image) && (
                <img className={classes.img} alt="img" src={model.imageUrl(props.blog.image, 300)} />
            )}
          </span>
          <span style={{color: 'var(--card-fg-color)'}}>{props.blog.overview}</span>
        </Typography>
      </CardContent>
      <CardActions>
        <Typography gutterBottom variant="overline" component="span">
          {strings.formatString(strings.edited_at, {
            time: formatter.format_time(props.blog.updated_at)
          })}
        </Typography>
      </CardActions>
    </Card>
  );
}
