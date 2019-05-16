import {
	FILE_UPLOAD_CHANGED,
	CATEGORY_UPLOAD_CHANGED,
	YEAR_UPLOAD_CHANGED,
	RANK_UPLOAD_SUCCESS,
	RANK_UPLOAD_LOADING,
	RANK_UPLOAD_FAIL,
	FORMAT_UPLOAD_CHANGED,
	PARIS_UPLOAD_CHANGED,
	RANK_UPLOAD_ERROR_RESET,
} from '../utils/Types';

const INITIAL_STATE = {
	file: null,
	year: '',
	category: '',
	format: '',
	isParis: false,
	// TODO: move somewhere where this makes sense
	loading: false,
	error: '',
};

export default (state = INITIAL_STATE, action) => {
	switch (action.type) {
		case FILE_UPLOAD_CHANGED:
			return {...state, error: '', file: action.payload};
		case CATEGORY_UPLOAD_CHANGED:
			return {...state, error: '', category: action.payload};
		case YEAR_UPLOAD_CHANGED:
			return {...state, error: '', year: action.payload};
		case RANK_UPLOAD_SUCCESS:
			return {...INITIAL_STATE, error: '', loading: false};
		case RANK_UPLOAD_LOADING:
			return {...state, error: '', loading: true};
		case RANK_UPLOAD_FAIL:
			return {...state, error: action.payload, loading: false};
		case FORMAT_UPLOAD_CHANGED:
			return {...state, error: '', format: action.payload};
		case PARIS_UPLOAD_CHANGED:
			return {...state, error: '', isParis: action.payload};
		case RANK_UPLOAD_ERROR_RESET:
			return {...state, error: ''};
		default:
			return state;
	}
};
