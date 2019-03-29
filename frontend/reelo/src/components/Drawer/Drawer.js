import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import BarApp from '../MyAppBar';
import DrawerList from './DrawerList';

const drawerWidth = '13%';

const styles = theme => ({
	root: {
		display: 'flex',
		flexGrow: 1,
	},
	menuButton: {
		marginLeft: -12,
		marginRight: 20,
	},
	spazio: {
		marginTop: 64,
		marginLeft: 'auto',
		marginRight: 'auto',
	},
	drawer: {
		width: drawerWidth,
		flexShrink: 0,
	},
	drawerPaper: {
		width: drawerWidth,
		backgroundColor: theme.palette.secondary.main,
	},
	content: {
		flexGrow: 1,
		padding: theme.spacing.unit * 3,
	},
	toolbar: theme.mixins.toolbar,
});

function ClippedDrawer(props) {
	/**
	 * 2019/03/04 created Gerardo
	 */

	const {classes} = props;

	return (
		<div className={classes.root}>
			<BarApp />
			<Drawer
				className={classes.drawer}
				variant="permanent"
				classes={{
					paper: classes.drawerPaper,
				}}>
				<div className={classes.toolbar} />
				<DrawerList />
			</Drawer>
			<div className={classes.spazio}>{props.children}</div>
		</div>
	);
}

ClippedDrawer.propTypes = {
	classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ClippedDrawer);
