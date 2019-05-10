import {combineReducers} from 'redux';
import loginReducer from './loginReducer';
import authReducer from './authReducer';
import uploadReducer from './uploadReducer';

export default combineReducers({
	loginForm: loginReducer,
	auth: authReducer,
	uploadForm: uploadReducer,
});
