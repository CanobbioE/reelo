import axios from "axios";
import {
    NAMESAKE_FETCH_LOADING,
    NAMESAKE_FETCH_SUCCESS,
    NAMESAKE_POST_LOADING,
    NAMESAKE_POST_SUCCESS,
    NAMESAKE_COMMENT_LOADING,
    NAMESAKE_COMMENT_SUCCESS,
    NAMESAKE_UPDATE,
    ERROR,
    ERROR_RESET,
} from "../utils/Types";
import Globals from "../config/Globals";

export const fetchNamesakes = (page = 1, size = -1) => async dispatch => {
    dispatch({
        type: NAMESAKE_FETCH_LOADING,
    });
    try {
        const response = await axios.get(
            `${Globals.baseURL}${Globals.API.namesakes.all}?page=${page}&size=${size}`,
            {
                headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
            },
        );
        dispatch({
            type: NAMESAKE_FETCH_SUCCESS,
            payload: response.data,
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};

export const updateNamesake = (index, namesake) => ({
    type: NAMESAKE_UPDATE,
    payload: { index, namesake },
});

export const acceptNamesake = namesake => async dispatch => {
    dispatch({
        type: NAMESAKE_POST_LOADING,
    });
    dispatch({
        type: ERROR_RESET,
    });
    try {
        await axios.post(`${Globals.baseURL}${Globals.API.namesakes.update}`, namesake, {
            headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
        });
        dispatch({
            type: NAMESAKE_POST_SUCCESS,
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};

export const commentNamesake = (namesake, comment) => async dispatch => {
    dispatch({
        type: NAMESAKE_COMMENT_LOADING,
    });
    dispatch({
        type: ERROR_RESET,
    });
    try {
        await axios.post(
            `${Globals.baseURL}${Globals.API.players.comment}`,
            {
                text: comment,
                player: { ...namesake.player },
            },
            {
                headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
            },
        );
        dispatch({
            type: NAMESAKE_COMMENT_SUCCESS,
        });
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};
