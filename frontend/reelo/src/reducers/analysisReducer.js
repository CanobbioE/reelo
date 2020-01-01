import {
	NAMESAKE_FETCH_LOADING,
	NAMESAKE_FETCH_ERROR,
	NAMESAKE_FETCH_SUCCESS,
	NAMESAKE_UPDATE,
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
		case NAMESAKE_UPDATE:
			return {
				...state,
				error: '',
				namesakes: updateSolvers(
					action.payload.index,
					action.payload.namesake,
					state.namesakes,
				),
			};
		default:
			return {...state};
	}
};

const updateSolvers = (index, namesake, namesakes) => {
	let tmp = namesakes.slice(0, index + 1);
	tmp[index] = namesake;
	tmp = tmp.concat(namesakes.slice(index + 2));
	return tmp;
};
