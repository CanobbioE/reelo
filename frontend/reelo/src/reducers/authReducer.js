import {AUTH_USER, AUTH_ERROR} from '../utils/Types';

const INITIAL_STATE = {
	authenticated:
		localStorage.getItem('token') !== '' && localStorage.getItem('token'),
	errorMessage: [],
};

export default (state = INITIAL_STATE, action) => {
	switch (action.type) {
		case AUTH_USER:
			return {
				...state,
				authenticated: action.payload,
				errorMessage: [],
			};
		case AUTH_ERROR:
			return {...state, errorMessage: action.payload};
		default:
			return state;
	}
};