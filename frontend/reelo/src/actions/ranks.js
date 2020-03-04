import axios from "axios";
import {
    RANKS_FETCH_LOADING,
    RANKS_FETCH_SUCCESS,
    RANKS_PAGE_SET,
    RANKS_COUNT_LOADING,
    RANKS_SIZE_SET,
    RANKS_COUNT_SUCCESS,
    RANKS_YEARS_LOADING,
    RANKS_YEARS_SUCCESS,
    ERROR_RESET,
    ERROR,
} from "../utils/Types";
import Globals from "../config/Globals";

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
    dispatch({
        type: ERROR_RESET,
    });
    try {
        const response = await axios.get(`${Globals.baseURL}${Globals.API.players.count}`);
        dispatch({
            type: RANKS_COUNT_SUCCESS,
            payload: response.data,
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};

export const fetchAllYears = () => async dispatch => {
    dispatch({
        type: RANKS_YEARS_LOADING,
    });
    dispatch({
        type: ERROR_RESET,
    });
    try {
        const response = await axios.get(`${Globals.baseURL}${Globals.API.ranks.years}`);
        dispatch({
            type: RANKS_YEARS_SUCCESS,
            payload: response.data,
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};

export const fetchRanks = (page = 1, size = -1) => async dispatch => {
    dispatch({
        type: RANKS_FETCH_LOADING,
    });
    dispatch({
        type: ERROR_RESET,
    });
    try {
        const response = await axios.get(
            `${Globals.baseURL}${Globals.API.ranks.all}/?page=${page}&size=${size}`,
            {
                headers: {
                    Authorization: localStorage.getItem("token"),
                },
            },
        );
        dispatch({
            type: RANKS_FETCH_SUCCESS,
            payload: response.data,
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};

export const forceReelo = () => async dispatch => {
    dispatch({
        type: RANKS_FETCH_LOADING,
    });
    dispatch({
        type: ERROR_RESET,
    });
    try {
        await axios.put(`${Globals.baseURL}${Globals.API.players.reelo.calculate}`, null, {
            headers: {
                Authorization: localStorage.getItem("token"),
            },
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};
