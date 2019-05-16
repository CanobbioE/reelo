import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import MyAppBar from '../MyAppBar';
import DrawerList from './DrawerList';
import {connect} from 'react-redux';

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
			<MyAppBar />
			<Drawer
				className={classes.drawer}
				variant="permanent"
				classes={{
					paper: classes.drawerPaper,
				}}>
				<div className={classes.toolbar} />
				<DrawerList isAuthenticated={props.auth.authenticated} />
			</Drawer>
			<div className={classes.spacing}>{props.children}</div>
		</div>
	);
}

ClippedDrawer.propTypes = {
	classes: PropTypes.object.isRequired,
};

// Using redux in this component is ok because it is actually more a container
function mapStateToProps({auth}) {
	return {auth};
}

const composedComponent = connect(mapStateToProps);

export default withStyles(styles)(composedComponent(ClippedDrawer));
