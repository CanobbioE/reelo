import React from 'react';
import {connect} from 'react-redux';
import {withStyles} from '@material-ui/core/styles';
import {AppBar, Toolbar, Typography, Button} from '@material-ui/core';
import {Link} from 'react-router-dom';
import Globals from '../config/Globals';
import {signout} from '../actions';

const styles = theme => ({
	root: {
		flexGrow: 1,
		marginBottom: '10px',
		zIndex: theme.zIndex.drawer + 1,
		backgroundColor: theme.palette.primary,
	},
	grow: {
		flexGrow: 1,
	},
	menuButton: {
		right: '24px',
		position: 'absolute',
		display: 'flex',
	},
	appname: {
		textDecoration: 'none',
	},
	navbtn: {
		marginRight: theme.spacing.unit * 3,
	},
});

function NavBar(props) {
	const {classes} = props;

	const logBtn = props.auth.authenticated ? (
		<Button
			component={Link}
			to={Globals.routes.home}
			onClick={props.signout}
			color="secondary">
			<Typography color="secondary" variant="subtitle1">
				Esci
			</Typography>
		</Button>
	) : null;

	return (
		<div className={classes.root}>
			<AppBar position="static" color="primary">
				<Toolbar>
					<Button
						className={classes.navbtn}
						component={Link}
						to={Globals.routes.home}
						color="inherit">
						<Typography color="secondary" variant="subtitle1">
							Classifiche
						</Typography>
					</Button>

					{props.auth.authenticated && (
						<Button
							className={classes.navbtn}
							component={Link}
							to={Globals.routes.upload}
							color="inherit">
							<Typography color="secondary" variant="subtitle1">
								Carica Classifiche
							</Typography>
						</Button>
					)}

					{props.auth.authenticated && (
						<Button
							className={classes.navbtn}
							component={Link}
							to={Globals.routes.varchange}
							color="inherit">
							<Typography color="secondary" variant="subtitle1">
								Modifica Algoritmo
							</Typography>
						</Button>
					)}

					{props.auth.authenticated && (
						<Button
							className={classes.navbtn}
							component={Link}
							to={Globals.routes.namesakes}
							color="inherit">
							<Typography color="secondary" variant="subtitle1">
								Omonimi
							</Typography>
						</Button>
					)}

					<Button
						className={classes.navbtn}
						component={Link}
						to={Globals.routes.about}
						color="inherit">
						<Typography color="secondary" variant="subtitle1">
							Informazioni
						</Typography>
					</Button>

					<div className={classes.menuButton}>{logBtn}</div>
				</Toolbar>
			</AppBar>
		</div>
	);
}

function mapStateToProps({auth}) {
	return {auth};
}

const composedComponent = connect(
	mapStateToProps,
	{
		signout,
	},
);

export default withStyles(styles)(composedComponent(NavBar));
