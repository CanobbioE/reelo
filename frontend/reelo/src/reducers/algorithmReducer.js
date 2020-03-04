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
} from "../utils/Types";

const INITIAL_STATE = {
    year: "",
    ex: "",
    final: "",
    mult: "",
    exp: "",
    np: "",
    loading: false,
    currentValues: {},
};

export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case ALG_YEAR_CHANGED:
            return { ...state, year: action.payload };
        case ALG_EX_CHANGED:
            return { ...state, ex: action.payload };
        case ALG_FINAL_CHANGED:
            return { ...state, final: action.payload };
        case ALG_MULT_CHANGED:
            return { ...state, mult: action.payload };
        case ALG_EXP_CHANGED:
            return { ...state, exp: action.payload };
        case ALG_NP_CHANGED:
            return { ...state, np: action.payload };
        case ALG_FETCH_SUCCESS:
            return { ...state, currentValues: action.payload };
        case ALG_POST_SUCCESS:
            return { ...state, loading: false };
        case ALG_POST_LOADING:
            return { ...state, loading: true };
        default:
            return { ...state };
    }
};
