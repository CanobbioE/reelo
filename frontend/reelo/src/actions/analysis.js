import axios from 'axios';
import {
	RANKS_FETCH_ERROR,
	RANK_UPLOAD_ERROR_RESET,
	RANKS_FETCH_LOADING,
} from '../utils/Types';
import Globals from '../config/Globals';

export const purgePlayers = () => async dispatch => {
	dispatch({
		type: RANKS_FETCH_LOADING,
	});
	try {
		await axios.post(
			`${Globals.baseURL}${Globals.API.purge}`,
			{},
			{headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}},
		);
		dispatch({
			type: RANK_UPLOAD_ERROR_RESET,
		});
	} catch (e) {
		dispatch({
			type: RANKS_FETCH_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};
