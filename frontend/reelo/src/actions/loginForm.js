import axios from "axios";

import {
    EMAIL_SIGNIN_CHANGED,
    PASSWORD_SIGNIN_CHANGED,
    SIGNIN_FORM_RESET,
    AUTH_USER,
    ERROR,
    ERROR_RESET,
} from "../utils/Types";
import Globals from "../config/Globals";

export const updateEmail = email => {
    return {
        type: EMAIL_SIGNIN_CHANGED,
        payload: email,
    };
};

export const updatePassword = password => {
    return {
        type: PASSWORD_SIGNIN_CHANGED,
        payload: password,
    };
};

export const signin = (email, password) => async dispatch => {
    dispatch({
        type: ERROR_RESET,
    });
    try {
        const response = await axios.post(`${Globals.baseURL}${Globals.API.auth.login}`, {
            email,
            password,
        });
        localStorage.setItem("token", response.data);
        dispatch({
            type: AUTH_USER,
            payload: response.data,
        });
        dispatch({
            type: SIGNIN_FORM_RESET,
        });
    } catch (e) {
        console.log(e.response.data);
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || "server offline",
        });
        dispatch({
            type: SIGNIN_FORM_RESET,
        });
    }
};

export const signout = () => dispatch => {
    localStorage.removeItem("token");
    dispatch({
        type: AUTH_USER,
        payload: "",
    });
};
