import {
	NAMESAKE_FETCH_LOADING,
	NAMESAKE_FETCH_ERROR,
	NAMESAKE_FETCH_SUCCESS,
} from '../utils/Types';

const INITIAL_STATE = {
	loading: false,
	error: '',
	namesakes: [],
	fixedNamesakes: {},
};

export default (state = INITIAL_STATE, action) => {
	switch (action.type) {
		case NAMESAKE_FETCH_ERROR:
			return {...state, error: action.payload, loading: false};
		case NAMESAKE_FETCH_LOADING:
			return {...state, error: '', namesakes: [], loading: true};
		case NAMESAKE_FETCH_SUCCESS:
			return {...state, namesakes: action.payload, loading: false};
		default:
			return {...state};
	}
};
