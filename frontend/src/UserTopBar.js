import React from 'react';
import {withStyles} from '@material-ui/core/styles';
import BooSticky from './BooSticky';
import model from './model';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import ButtonBase from '@material-ui/core/ButtonBase';

const styles = (theme => ({
	root: {
		backgroundColor: 'var(--card-bg-color)',
		color: 'var(--card-fg-color)',
    boxShadow: '1px 1px 20px var(--shadow-color)',
    display: 'block'
	},
  a: {
    color: 'var(--blog-a-color)',
    textDecoration: 'none',
    '&:visited': {
      color: 'var(--blog-a-color)',
      textDecoration: 'none',
    }
  },
  label: {
    marginRight: '5px',
    color: 'var(--card-fg2-color)'
  },
	user: {
    padding: '40px',
    [theme.breakpoints.down('xs')]: {
      padding: '20px 30px'
    }
	},
	sticky: {
		width: '100%',
		zIndex: 100,
		background: "white"
	},
	stickyRaised: {
		width: '100%',
		zIndex: 100,
		background: 'white',
		boxShadow: '0px 0px 5px var(--shadow-color)'
	},
	img: {
		width: '80px',
		height: '80px',
		borderRadius: '80px',
    [theme.breakpoints.down('xs')]: {
      width: '40px',
      height: '40px',
      borderRadius: '40px',
    }
	}
}));

class UserTopBar extends React.Component {
  constructor(props) {
    super(props);
    const {classes} = this.props;
    this.state = {
      stickyClass: classes.sticky,
      disableSticky: false,
      info: {
        name: 'USER NAME',
        bio: 'has no bio info',
        avatar: 'https://p1.ssl.qhmsg.com/dm/200_200_80/t012880d8265009ca96.jpg',
        contacts: []
      }
    };
    this.sticky = React.createRef();
  }

  componentDidMount() {
    model.userInfo().then(info => {
      this.setState({info: info});
    });
    window.addEventListener('lang-changed', e => {
      model.userInfo().then(info => {
        this.setState({info: info});
      });
    });
  }

  onStickyRaised(raised) {
    const {classes} = this.props;
		if (raised) {
      this.setState({stickyClass: classes.stickyRaised});
		} else {
      this.setState({stickyClass: classes.sticky});
		}
  }

  render() {
    let { classes } = this.props;
    return (
      <div className={classes.root} ref="root">
        <Grid className={classes.user} container spacing={2}>
          <Grid item>
            <ButtonBase className={classes.image}>
              <img className={classes.img} alt="complex" src={this.state.info.avatar} />
            </ButtonBase>
          </Grid>
          <Grid item xs={12} sm container>
            <Grid item xs container direction="column" spacing={2}>
              <Grid item xs>
                <Typography gutterBottom variant="subtitle1">
                  {this.state.info.name}
                </Typography>
                <Typography variant="body2" gutterBottom>
                  {this.state.info.bio}
                </Typography>
                {this.state.info.contacts.map((contacts) => (
                  <Typography key={contacts.value} variant="body2">
                    <label className={classes.label}>{contacts.label}: </label>
                    {contacts.href ? 
                      <a className={classes.a} href={contacts.href}>{contacts.value}</a> :
                      <span style={{color: 'var(--card-fg2-color)'}}>{contacts.value}</span>}
                  </Typography>
                ))}
              </Grid>
            </Grid>
          </Grid>
        </Grid>
        <BooSticky ref={this.sticky} disabled={this.state.disableSticky} top={0} onRaised={this.onStickyRaised.bind(this)}>
          <div className={this.state.stickyClass}>
            <div style={{maxWidth: '100%'}}>{this.props.children}</div>
          </div>
        </BooSticky>
      </div>
    );
  }

  root() {
    return this.refs["root"];
  }

  disableSticky(disable) {
    this.setState({disableSticky: disable});
  }

  updateSticky() {
    this.sticky.current.updateSticky();
  }
}

export default withStyles(styles)(UserTopBar);
