import {Grid} from '@material-ui/core';
import React from 'react';
import LoginForm from '../components/LoginForm';
import Logout from '../components/Logout';
import {updateEmail, updatePassword, signin, signout} from '../actions';
import {connect} from 'react-redux';
import Globals from '../config/Globals';

const Admin = props => {
	const login = event => {
		event.preventDefault();
		props.signin(props.loginForm.email, props.loginForm.password);
		props.history.push(Globals.routes.home);
	};

	const logout = event => {
		event.preventDefault();
		props.signout();
		props.history.push(Globals.routes.home);
	};

	const loginForm = (
		<LoginForm
			onPasswordChange={props.updatePassword}
			onEmailChange={props.updateEmail}
			onSubmit={login}
			emailValue={props.loginForm.email}
			passwordValue={props.loginForm.password}
		/>
	);
	const logoutForm = <Logout onClick={logout} />;
	const form = props.auth.authenticated ? logoutForm : loginForm;

	return (
		<Grid container justify="center">
			{form}
		</Grid>
	);
};

function mapStateToProps({loginForm, auth}) {
	return {loginForm, auth};
}

const composedComponent = connect(
	mapStateToProps,
	{
		updateEmail,
		updatePassword,
		signin,
		signout,
	},
);

export default composedComponent(Admin);
