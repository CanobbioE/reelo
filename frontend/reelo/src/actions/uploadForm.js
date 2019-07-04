import axios from 'axios';
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
	START_UPLOAD_CHANGED,
	END_UPLOAD_CHANGED,
} from '../utils/Types';
import Globals from '../config/Globals';

export const updateUploadFile = file => {
	return {
		type: FILE_UPLOAD_CHANGED,
		payload: file,
	};
};

export const resetUploadForm = () => {
	return {
		type: RANK_UPLOAD_ERROR_RESET,
	};
};

export const updateUploadIsParis = isParis => {
	return {
		type: PARIS_UPLOAD_CHANGED,
		payload: isParis,
	};
};

export const updateUploadYear = year => {
	return {
		type: YEAR_UPLOAD_CHANGED,
		payload: year,
	};
};

export const updateUploadCategory = category => {
	return {
		type: CATEGORY_UPLOAD_CHANGED,
		payload: category,
	};
};

export const updateUploadFormat = format => {
	return {
		type: FORMAT_UPLOAD_CHANGED,
		payload: format,
	};
};

export const updateUploadStart = start => {
	return {
		type: START_UPLOAD_CHANGED,
		payload: start,
	};
};

export const updateUploadEnd = end => {
	return {
		type: END_UPLOAD_CHANGED,
		payload: end,
	};
};

const fieldConverter = field => {
	field.trim();
	switch (field) {
		case 'n':
			return 'nome';
		case 'c':
			return 'cognome';
		case 's':
		case 'citta':
		case 'sede':
			return 'città';
		case 'p':
		case 'punteggio':
			return 'punti';
		case 't':
			return 'tempo';
		case 'e':
		case 'es':
			return 'esercizi';
		case ' ':
			return '';
		default:
			return field;
	}
};

export const uploadFile = (
	file,
	category,
	year,
	isParis,
	format,
	start,
	end,
) => async dispatch => {
	dispatch({
		type: RANK_UPLOAD_LOADING,
	});
	try {
		const jwt = localStorage.getItem('token');
		const mappedFormat = format
			.replace(',', ' ')
			.split(' ')
			.map(field => fieldConverter(field.toLowerCase()))
			.reduce((f1, f2) => f1 + ' ' + f2);
		const data = JSON.stringify({
			category: category,
			year: year,
			isParis: isParis,
			token: jwt,
			format: mappedFormat,
			start: start,
			end: end,
		});
		const formData = new FormData();
		formData.append('file', file);
		formData.append('data', data);
		await axios.post(`${Globals.baseURL}${Globals.API.upload}`, formData, {
			headers: {
				'Content-Type': 'multipart/form-data',
				Authorization: `Bearer ${localStorage.getItem('token')}`,
			},
		});

		dispatch({
			type: RANK_UPLOAD_SUCCESS,
		});
		// TODO: This is horrible, this is madness
		alert('Caricamento avvenuto con successo');
	} catch (e) {
		console.log(e);
		dispatch({
			type: RANK_UPLOAD_FAIL,
			payload: e.response.data,
		});
	}
};