import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import BarApp from '../MyAppBar';
import DrawerList from './DrawerList';

const drawerWidth = '15%';

const styles = theme => ({
	root: {
		display: 'flex',
		flexGrow: 1,
	},
	spacing: {
		marginTop: '8%',
		width: '100%',
	},
	drawer: {
		width: drawerWidth,
		flexShrink: 0,
	},
	drawerPaper: {
		width: drawerWidth,
		backgroundColor: theme.palette.secondary.main,
	},
	toolbar: theme.mixins.toolbar,
});

function ClippedDrawer(props) {
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
			<div className={classes.spacing}>{props.children}</div>
		</div>
	);
}

ClippedDrawer.propTypes = {
	classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ClippedDrawer);
