import axios from "axios";
import {
    ALG_YEAR_CHANGED,
    ALG_EX_CHANGED,
    ALG_FINAL_CHANGED,
    ALG_MULT_CHANGED,
    ALG_EXP_CHANGED,
    ALG_NP_CHANGED,
    ALG_FETCH_SUCCESS,
    ALG_POST_SUCCESS,
    ALG_POST_LOADING,
    ALG_POST_FAIL,
} from "../utils/Types";
import Globals from "../config/Globals";

export const updateAlgYear = year => {
    return {
        type: ALG_YEAR_CHANGED,
        payload: year,
    };
};
export const updateAlgEx = ex => {
    return {
        type: ALG_EX_CHANGED,
        payload: ex,
    };
};
export const updateAlgFinal = final => {
    return {
        type: ALG_FINAL_CHANGED,
        payload: final,
    };
};
export const updateAlgMult = mult => {
    return {
        type: ALG_MULT_CHANGED,
        payload: mult,
    };
};
export const updateAlgExp = exp => {
    return {
        type: ALG_EXP_CHANGED,
        payload: exp,
    };
};
export const updateAlgNP = np => {
    return {
        type: ALG_NP_CHANGED,
        payload: np,
    };
};
export const fetchVars = () => async dispatch => {
    try {
        const response = await axios.get(`${Globals.baseURL}${Globals.API.costants.all}`, {
            headers: {
                Authorization: localStorage.getItem("token"),
            },
        });
        dispatch({
            type: ALG_FETCH_SUCCESS,
            payload: response.data,
        });
    } catch (e) {
        console.log(e);
    }
};
export const updateAlg = (year, ex, final, mult, exp, np, curr) => async dispatch => {
    dispatch({
        type: ALG_POST_LOADING,
    });
    try {
        // TODO: OMG this is orrible please use Object.keys
        if (year === "") year = curr.year;
        if (ex === "") ex = curr.ex;
        if (final === "") final = curr.final;
        if (mult === "") mult = curr.mult;
        if (exp === "") exp = curr.exp;
        if (np === "") np = curr.np;
        console.log("aaaaaaaaaaaaaa", year, ex, final, mult, exp, np)
        const v = {
            year: parseInt(`${year}`.replace(",", ".")),
            ex: parseFloat(`${ex}`.replace(",", ".")),
            final: parseFloat(`${final}`.replace(",", ".")),
            mult: parseFloat(`${mult}`.replace(",", ".")),
            exp: parseFloat(`${exp}`.replace(",", ".")),
            np: parseFloat(`${np}`.replace(",", ".")),
        };
        await axios.patch(`${Globals.baseURL}${Globals.API.costants.update}`, v, {
            headers: {
                "Content-Type": "multipart/form-data",
                Authorization: localStorage.getItem("token"),
            },
        });
        dispatch({
            type: ALG_POST_SUCCESS,
        });
    } catch (e) {
        console.log(e);
        dispatch({
            type: ALG_POST_FAIL,
            payload: (e && e.response && e.response.data) || "Errore inaspettato",
        });
    }
};
