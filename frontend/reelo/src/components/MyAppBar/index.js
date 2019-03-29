import React from 'react';
import {AppBar, withStyles, Toolbar, Typography} from '@material-ui/core';
import {properties} from '../../config/Properties';

const styles = theme => ({
	appBar: {
		zIndex: theme.zIndex.drawer + 1,
		backgroundColor: theme.palette.primary,
	},
	grow: {
		flexGrow: 1,
	},
});

function BarApp(props) {
	const {classes} = props;
	return (
		<AppBar position="fixed" className={classes.appBar}>
			<Toolbar>
				<Typography variant="h4" color="inherit" className={classes.grow}>
					{properties.appname}
				</Typography>
			</Toolbar>
		</AppBar>
	);
}
export default withStyles(styles)(BarApp);
