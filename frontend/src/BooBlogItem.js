import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles(theme =>({
  card: {
    borderRadius: '0px',
    boxShadow: 'none'
  },
  area: {
    '&:hover': {
      backgroundColor: 'white'
    }
  },
  img: {
    maxHeight: '100%',
    width: 'auto',
    borderRadius: '2px',
    maxWidth: '40%',
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
      <CardActionArea>
        <CardContent className={classes.area} onClick={() => window.boo.location.go("/" + props.blog.url_id)}>
          <Typography gutterBottom variant="h5" component="h2">
            {props.blog.title}
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            {Boolean(props.blog.image_url) && (
              <img className={classes.img} alt="img" align="left" hspace="10" vspace="3" src={props.blog.image_url} />
            )}
            {props.blog.overview}
          </Typography>
        </CardContent>
        <CardActions>
          <Typography gutterBottom variant="overline" component="span">
            发表于3天前
          </Typography>
          ,
          <Typography gutterBottom variant="overline" component="span">
            1天前修改过
          </Typography>
        </CardActions>
      </CardActionArea>
    </Card>
  );
}
