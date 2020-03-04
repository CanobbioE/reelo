import axios from "axios";
import {
    FILE_UPLOAD_CHANGED,
    CATEGORY_UPLOAD_CHANGED,
    YEAR_UPLOAD_CHANGED,
    RANK_UPLOAD_SUCCESS,
    RANK_UPLOAD_LOADING,
    FORMAT_UPLOAD_CHANGED,
    PARIS_UPLOAD_CHANGED,
    START_UPLOAD_CHANGED,
    END_UPLOAD_CHANGED,
    ERROR_RESET,
    ERROR,
} from "../utils/Types";
import Globals from "../config/Globals";

export const checkExistence = (year, category, isParis) => async dispatch => {
    dispatch({
        type: RANK_UPLOAD_LOADING,
    });
    try {
        const response = await axios.get(
            `${Globals.baseURL}${Globals.API.ranks.exist}/?y=${year}&cat=${category}&isparis=${isParis}`,
            {
                headers: {
                    Authorization: localStorage.getItem("token"),
                },
            },
        );
        dispatch({
            type: ERROR_RESET,
        });
        return response.data;
    } catch (e) {
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};

export const updateUploadFile = file => {
    return {
        type: FILE_UPLOAD_CHANGED,
        payload: file,
    };
};

export const resetUploadForm = () => {
    return {
        type: ERROR_RESET,
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
        payload: year.trim(),
    };
};

export const updateUploadCategory = category => {
    return {
        type: CATEGORY_UPLOAD_CHANGED,
        payload: category.trim(),
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
        payload: start.trim(),
    };
};

export const updateUploadEnd = end => {
    return {
        type: END_UPLOAD_CHANGED,
        payload: end.trim(),
    };
};

const fieldConverter = field => {
    field.trim();
    switch (field) {
        case "n":
            return "nome";
        case "c":
            return "cognome";
        case "s":
        case "citta":
        case "sede":
            return "città";
        case "p":
        case "punteggio":
            return "punti";
        case "t":
            return "tempo";
        case "e":
        case "es":
            return "esercizi";
        case " ":
            return "";
        default:
            return field;
    }
};

export const uploadFile = (file, category, year, isParis, format, start, end) => async dispatch => {
    dispatch({
        type: RANK_UPLOAD_LOADING,
    });
    dispatch({
        type: ERROR_RESET,
    });
    try {
        const jwt = localStorage.getItem("token");
        const mappedFormat = format
            .replace(",", " ")
            .split(" ")
            .map(field => fieldConverter(field.toLowerCase()))
            .reduce((f1, f2) => f1 + " " + f2);
        const data = JSON.stringify({
            game: {
                category: category,
                year: parseInt(year),
                isParis: isParis,
                start: parseInt(start),
                end: parseInt(end),
            },
            token: jwt,
            format: mappedFormat.trim(),
        });
        const formData = new FormData();
        formData.append("file", file);
        formData.append("data", data);
        await axios.post(`${Globals.baseURL}${Globals.API.ranks.upload}`, formData, {
            headers: {
                "Content-Type": "multipart/form-data",
                Authorization: `Bearer ${localStorage.getItem("token")}`,
            },
        });

        dispatch({
            type: RANK_UPLOAD_SUCCESS,
        });
        // TODO: This is horrible, this is madness
        alert("Caricamento avvenuto con successo");
    } catch (e) {
        console.log(e);
        dispatch({
            type: ERROR,
            payload: (e && e.response && e.response.data) || {code: "NO_CONN"},
        });
    }
};
