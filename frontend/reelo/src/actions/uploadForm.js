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
} from '../utils/Types';
import Globals from '../config/Globals';

export const updateUploadFile = file => {
	return {
		type: FILE_UPLOAD_CHANGED,
		payload: file,
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

export const uploadFile = (
	file,
	category,
	year,
	isParis,
	format,
) => async dispatch => {
	var response;
	dispatch({
		type: RANK_UPLOAD_LOADING,
	});
	try {
		const jwt = localStorage.getItem('token');
		const data = JSON.stringify({
			category: category,
			year: year,
			isParis: isParis,
			token: jwt,
			format: format,
		});
		const formData = new FormData();
		formData.append('file', file);
		formData.append('data', data);
		response = await axios.post(
			`${Globals.baseURL}${Globals.API.upload}`,
			formData,
			{
				headers: {
					'Content-Type': 'multipart/form-data',
				},
			},
		);
		dispatch({
			type: RANK_UPLOAD_SUCCESS,
		});
	} catch (e) {
		console.log(e);
		dispatch({
			type: RANK_UPLOAD_FAIL,
			payload: response,
		});
	}
};
