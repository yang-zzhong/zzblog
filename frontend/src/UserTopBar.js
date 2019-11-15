import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import BooSticky from './BooSticky';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import ButtonBase from '@material-ui/core/ButtonBase';

const useStyles = makeStyles(theme => ({
	root: {
		backgroundColor: 'white',
		borderBottomWidth: '1px',
		borderBottomStyle: 'solid',
		borderBottomColor: '#f0f0f0',
		color: theme.palette.secondary
	},
	user: {
    padding: '40px',
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
		boxShadow: '0px 0px 2px rgba(0, 0, 0, .4)'
	},
	img: {
		width: '80px',
		height: '80px',
		borderRadius: '80px'
	}
}));

export default function UserTopBar(props) {
	let classes = useStyles();
	const [stickyClass, setStickyClass] = React.useState(null);
	const onStickyRaised = raised => {
		if (raised) {
			setStickyClass(classes.stickyRaised);
		} else {
			setStickyClass(classes.sticky);
		}
		if (props.children) {
			window.dispatchEvent(new CustomEvent('header-noshadow', {detail: raised}));
		}
	};

	return (
		<div className={classes.root}>
      <Grid className={classes.user} container spacing={2}>
        <Grid item>
          <ButtonBase className={classes.image}>
            <img className={classes.img} alt="complex" src="https://p1.ssl.qhmsg.com/dm/200_200_80/t012880d8265009ca96.jpg" />
          </ButtonBase>
        </Grid>
        <Grid item xs={12} sm container>
          <Grid item xs container direction="column" spacing={2}>
            <Grid item xs>
              <Typography gutterBottom variant="subtitle1">
								Hackyoung
              </Typography>
              <Typography variant="body2" gutterBottom>
								现在的仙女座星系到底是什么样子，必须要 250 万年之后我们才能看到。它是什么时候形成的？怎么形成的？如何演化的？将来会变成什么样？” 
              </Typography>
              <Typography variant="body2" color="textSecondary">
								hackyoung110@gmail.com
              </Typography>
            </Grid>
          </Grid>
        </Grid>
      </Grid>
			<BooSticky top={0} onRaised={onStickyRaised}>
				<div className={stickyClass}>
				  <div style={{maxWidth: '100%'}}>{props.children}</div>
				</div>
			</BooSticky>
		</div>
	);
}
