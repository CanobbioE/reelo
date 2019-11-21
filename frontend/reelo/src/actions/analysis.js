import axios from 'axios';
import {
	NAMESAKE_FETCH_LOADING,
	NAMESAKE_FETCH_ERROR,
	NAMESAKE_FETCH_SUCCESS,
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
