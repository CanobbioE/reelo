import {
	RANKS_FETCH_LOADING,
	RANKS_FETCH_SUCCESS,
	RANKS_FETCH_ERROR,
} from '../utils/Types';

const INITIAL_STATE = {
	loading: false,
	rows: [],
	error: '',
};

export default (state = INITIAL_STATE, action) => {
	switch (action.type) {
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
			name: rank.name,
			surname: rank.surname,
			category: rank.category,
			reelo: rank.reelo,
		});
	});

	const sorted = rows.sort((a, b) => b.reelo - a.reelo);
	return sorted;
};
