import axios from 'axios';
import {
	RANKS_FETCH_LOADING,
	RANKS_FETCH_SUCCESS,
	RANKS_FETCH_ERROR,
} from '../utils/Types';
import Globals from '../config/Globals';

export const fetchRanks = () => async dispatch => {
	dispatch({
		type: RANKS_FETCH_LOADING,
	});
	try {
		const response = await axios.get(`${Globals.baseURL}${Globals.API.ranks}`, {
			headers: {
				Authorization: localStorage.getItem('token'),
			},
		});
		dispatch({
			type: RANKS_FETCH_SUCCESS,
			payload: response.data,
		});
	} catch (e) {
		dispatch({
			type: RANKS_FETCH_ERROR,
			payload: e.response.data,
		});
	}
};

export const forceReelo = () => async dispatch => {
	dispatch({
		type: RANKS_FETCH_LOADING,
	});
	try {
		await axios.put(`${Globals.baseURL}${Globals.API.force}`, {
			headers: {
				Authorization: localStorage.getItem('token'),
			},
		});
	} catch (e) {
		dispatch({
			type: RANKS_FETCH_ERROR,
			payload: e.response.data,
		});
	}
};
