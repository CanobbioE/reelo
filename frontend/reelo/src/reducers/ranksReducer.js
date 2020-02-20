import {
	RANKS_FETCH_LOADING,
	RANKS_FETCH_SUCCESS,
	RANKS_FETCH_ERROR,
	RANKS_PAGE_SET,
	RANKS_COUNT_LOADING,
	RANKS_COUNT_ERROR,
	RANKS_COUNT_SUCCESS,
	RANKS_SIZE_SET,
	RANKS_YEARS_LOADING,
	RANKS_YEARS_ERROR,
	RANKS_YEARS_SUCCESS,
} from '../utils/Types';

const INITIAL_STATE = {
	loading: false,
	rows: [],
	error: '',
	page: 1,
	count: 0,
	size: 10,
	years: [],
};

export default (state = INITIAL_STATE, action) => {
	console.log(action)
	switch (action.type) {
		case RANKS_YEARS_LOADING:
			return {...state, loading: true, error: ''};
		case RANKS_YEARS_ERROR:
			return {...state, loading: false, error: action.payload};
		case RANKS_YEARS_SUCCESS:
			return {...state, years: action.payload, loading: false};
		case RANKS_FETCH_LOADING:
			return {...state, loading: true};
		case RANKS_FETCH_SUCCESS:
			return {
				...state,
				rows: ranksFetched(action.payload),
				loading: false,
				error: '',
			};
		case RANKS_FETCH_ERROR:
			return {...state, loading: false, error: action.payload};
		case RANKS_PAGE_SET:
			return {...state, page: action.payload};
		case RANKS_COUNT_LOADING:
			return {...state, loading: true, error: ''};
		case RANKS_COUNT_ERROR:
			return {...state, loading: false, error: action.payload};
		case RANKS_COUNT_SUCCESS:
			return {...state, count: action.payload, loading: false};
		case RANKS_SIZE_SET:
			return {...state, size: action.payload};
		default:
			return state;
	}
};

const ranksFetched = data => {
	const rows = [];
	if (!data || !data.length) return null;
	data.forEach((rank, index) => {
		rows.push({
			id: index,
			name: rank.player.name,
			surname: rank.player.surname,
			category: rank.lastCategory,
			reelo: rank.player.reelo,
			history: rank.history,
		});
	});

	const sorted = rows.sort((a, b) => b.reelo - a.reelo);
	return sorted;
};
