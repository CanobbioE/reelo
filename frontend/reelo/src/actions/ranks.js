import axios from 'axios';
import {
	RANKS_FETCH_LOADING,
	RANKS_FETCH_SUCCESS,
	RANKS_FETCH_ERROR,
	RANKS_PAGE_SET,
	RANKS_COUNT_LOADING,
	RANKS_COUNT_ERROR,
	RANKS_SIZE_SET,
	RANKS_COUNT_SUCCESS,
	RANKS_YEARS_LOADING,
	RANKS_YEARS_ERROR,
	RANKS_YEARS_SUCCESS,
} from '../utils/Types';
import Globals from '../config/Globals';

export const setRankPage = page => {
	return {
		type: RANKS_PAGE_SET,
		payload: page,
	};
};

export const setRankSize = size => {
	return {
		type: RANKS_SIZE_SET,
		payload: size,
	};
};

export const fetchTotalRanks = () => async dispatch => {
	dispatch({
		type: RANKS_COUNT_LOADING,
	});
	try {
		const response = await axios.get(`${Globals.baseURL}${Globals.API.count}`);
		dispatch({
			type: RANKS_COUNT_SUCCESS,
			payload: response.data,
		});
	} catch (e) {
		dispatch({
			type: RANKS_COUNT_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};

export const fetchAllYears = () => async dispatch => {
	dispatch({
		type: RANKS_YEARS_LOADING,
	});
	try {
		const response = await axios.get(`${Globals.baseURL}${Globals.API.years}`);
		dispatch({
			type: RANKS_YEARS_SUCCESS,
			payload: response.data,
		});
	} catch (e) {
		dispatch({
			type: RANKS_YEARS_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};

export const fetchRanks = (page, size) => async dispatch => {
	dispatch({
		type: RANKS_FETCH_LOADING,
	});
	try {
		const response = await axios.get(
			`${Globals.baseURL}${Globals.API.ranks}/?page=${page}&size=${size}`,
			{
				headers: {
					Authorization: localStorage.getItem('token'),
				},
			},
		);
		dispatch({
			type: RANKS_FETCH_SUCCESS,
			payload: response.data,
		});
	} catch (e) {
		dispatch({
			type: RANKS_FETCH_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};

export const forceReelo = () => async dispatch => {
	dispatch({
		type: RANKS_FETCH_LOADING,
	});
	try {
		await axios.put(`${Globals.baseURL}${Globals.API.force}`, null, {
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
