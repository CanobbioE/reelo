import axios from 'axios';

import {
	EMAIL_SIGNIN_CHANGED,
	PASSWORD_SIGNIN_CHANGED,
	SIGNIN_FORM_RESET,
	AUTH_USER,
	AUTH_ERROR,
	AUTH_LOADING,
} from '../helpers/Types';
import Globals from '../config/Globals';

export const updateSigninEmail = email => {
	return {
		type: EMAIL_SIGNIN_CHANGED,
		payload: email,
	};
};

export const updatePassword = password => {
	return {
		type: PASSWORD_SIGNIN_CHANGED,
		payload: password,
	};
};

export const signin = (email, password, callback) => async dispatch => {
	try {
		const response = await axios.post(
			`${Globals.baseURL}/${Globals.API.login}`,
			{
				email,
				password,
			},
		);
		dispatch({
			type: AUTH_USER,
			payload: response.headers.authorization,
		});
		dispatch({
			type: SIGNIN_FORM_RESET,
		});
		localStorage.setItem('token', response.headers.authorization);
		callback();
	} catch (e) {
		dispatch({
			type: AUTH_ERROR,
			payload: e && e.response && e.response.data && e.response.data.messages,
		});
		dispatch({
			type: SIGNIN_FORM_RESET,
		});
	}
};

export const signout = () => dispatch => {
	localStorage.removeItem('token');
	dispatch({
		type: AUTH_USER,
		payload: '',
	});
	//callback();
};
