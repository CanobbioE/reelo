import {
	EMAIL_SIGNIN_CHANGED,
	PASSWORD_SIGNIN_CHANGED,
	SIGNIN_FORM_RESET,
} from '../utils/Types';

const INITIAL_STATE = {
	email: '',
	password: '',
};

export default (state = INITIAL_STATE, action) => {
	switch (action.type) {
		case EMAIL_SIGNIN_CHANGED:
			return {...state, email: action.payload};
		case PASSWORD_SIGNIN_CHANGED:
			return {...state, password: action.payload};
		case SIGNIN_FORM_RESET:
			return {...state, email: '', password: ''};
		default:
			return state;
	}
};
