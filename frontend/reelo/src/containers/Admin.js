import {Grid} from '@material-ui/core';
import React from 'react';
import LoginForm from '../components/LoginForm';
import {updateSigninEmail, updatePassword, signin, signout} from '../actions';
import {connect} from 'react-redux';

function Admin(props) {
	const login = event => {
		event.preventDefault();
		props.signin('admin-canna', 'citrosodina');
	};

	return (
		<Grid container justify="center">
			<LoginForm onSubmit={login} />
		</Grid>
	);
}

function mapStateToProps({loginForm}) {
	return {loginForm};
}

const composedComponent = connect(
	mapStateToProps,
	{
		updateSigninEmail,
		updatePassword,
		signin,
		signout,
	},
);

export default composedComponent(Admin);
