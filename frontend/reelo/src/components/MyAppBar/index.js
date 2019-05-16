import React from 'react';
import {Link} from 'react-router-dom';
import {AppBar, withStyles, Toolbar, Typography} from '@material-ui/core';
import {properties} from '../../config/Properties';
import Globals from '../../config/Globals';

const styles = theme => ({
	appBar: {
		zIndex: theme.zIndex.drawer + 1,
		backgroundColor: theme.palette.primary,
	},
	grow: {
		flexGrow: 1,
		textDecoration: 'none',
	},
});

function MyAppBar(props) {
	const {classes} = props;
	return (
		<AppBar position="fixed" className={classes.appBar}>
			<Toolbar>
				<Typography
					component={Link}
					to={Globals.routes.home}
					variant="h4"
					color="inherit"
					className={classes.grow}>
					{properties.appname}
				</Typography>
			</Toolbar>
		</AppBar>
	);
}
export default withStyles(styles)(MyAppBar);
