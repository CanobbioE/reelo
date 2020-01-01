import axios from 'axios';
import {
	NAMESAKE_FETCH_LOADING,
	NAMESAKE_FETCH_ERROR,
	NAMESAKE_FETCH_SUCCESS,
	NAMESAKE_POST_LOADING,
	NAMESAKE_POST_ERROR,
	NAMESAKE_POST_SUCCESS,
	NAMESAKE_COMMENT_LOADING,
	NAMESAKE_COMMENT_ERROR,
	NAMESAKE_COMMENT_SUCCESS,
	NAMESAKE_UPDATE,
} from '../utils/Types';
import Globals from '../config/Globals';

export const fetchNamesakes = (page = 1, size = -1) => async dispatch => {
	dispatch({
		type: NAMESAKE_FETCH_LOADING,
	});
	try {
		const start = Date.now();
		const response = await axios.get(
			`${Globals.baseURL}${Globals.API.namesakes}/?page=${page}&size=${size}`,
			{
				headers: {Authorization: `Bearer ${localStorage.getItem('token')}`},
			},
		);
		console.log(`Duration for ${size} players = ${Date.now() - start}ms `);
		dispatch({
			type: NAMESAKE_FETCH_SUCCESS,
			payload: response.data,
		});
	} catch (e) {
		dispatch({
			type: NAMESAKE_FETCH_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};

export const updateNamesake = (index, namesake) => ({
	type: NAMESAKE_UPDATE,
	payload: {index, namesake},
});

export const acceptNamesake = namesake => async dispatch => {
	dispatch({
		type: NAMESAKE_POST_LOADING,
	});
	try {
		await axios.post(`${Globals.baseURL}${Globals.API.namesakes}`, namesake, {
			headers: {Authorization: `Bearer ${localStorage.getItem('token')}`},
		});
		dispatch({
			type: NAMESAKE_POST_SUCCESS,
		});
	} catch (e) {
		dispatch({
			type: NAMESAKE_POST_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};

export const commentNamesake = (namesake, comment) => async dispatch => {
	dispatch({
		type: NAMESAKE_COMMENT_LOADING,
	});
	try {
		await axios.post(
			`${Globals.baseURL}${Globals.API.comment}`,
			{...namesake, comment},
			{
				headers: {Authorization: `Bearer ${localStorage.getItem('token')}`},
			},
		);
		dispatch({
			type: NAMESAKE_COMMENT_SUCCESS,
		});
	} catch (e) {
		dispatch({
			type: NAMESAKE_COMMENT_ERROR,
			payload: e && e.response && e.response.data,
		});
	}
};
