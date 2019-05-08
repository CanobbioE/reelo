import {Grid} from '@material-ui/core';
import React from 'react';
import LoginForm from '../components/LoginForm';
import {updateEmail, updatePassword, signin} from '../actions';
import {connect} from 'react-redux';
import Globals from '../config/Globals';

const Admin = props => {
	const login = event => {
		event.preventDefault();
		props.signin(props.loginForm.email, props.loginForm.password);
		props.history.push(Globals.routes.home);
	};

	return (
		<Grid container justify="center">
			<LoginForm
				onPasswordChange={props.updatePassword}
				onEmailChange={props.updateEmail}
				onSubmit={login}
				emailValue={props.loginForm.email}
				passwordValue={props.loginForm.password}
			/>
		</Grid>
	);
};

function mapStateToProps({loginForm}) {
	return {loginForm};
}

const composedComponent = connect(
	mapStateToProps,
	{
		updateEmail,
		updatePassword,
		signin,
	},
);

export default composedComponent(Admin);
