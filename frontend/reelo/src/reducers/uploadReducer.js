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
} from "../utils/Types";
import { scoreHelp, formatHelp } from "../utils/Helper";

const INITIAL_STATE = {
    file: null,
    year: "",
    category: "",
    format: "",
    isParis: false,
    start: "",
    end: "",
    startSugg: "",
    endSugg: "",
    formatSugg: "",
    // TODO: move somewhere where this makes sense
    loading: false,
};

export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case FILE_UPLOAD_CHANGED:
            return { ...state, file: action.payload };
        case CATEGORY_UPLOAD_CHANGED:
            return {
                ...state,
                category: action.payload,
                startSugg: suggestedValueFor("start", state.year, action.payload),
                endSugg: suggestedValueFor("end", state.year, action.payload),
            };
        case YEAR_UPLOAD_CHANGED:
            return {
                ...state,
                year: action.payload,
                startSugg: suggestedValueFor("start", action.payload, state.category),
                endSugg: suggestedValueFor("end", action.payload, state.category),
                formatSugg: suggestedFormat(action.payload),
            };
        case RANK_UPLOAD_SUCCESS:
            return { ...INITIAL_STATE };
        case RANK_UPLOAD_LOADING:
            return { ...state, loading: true };
        case FORMAT_UPLOAD_CHANGED:
            return { ...state, format: action.payload };
        case PARIS_UPLOAD_CHANGED:
            return { ...state, isParis: action.payload };
        case START_UPLOAD_CHANGED:
            return { ...state, start: action.payload };
        case END_UPLOAD_CHANGED:
            return { ...state, end: action.payload };
        default:
            return state;
    }
};

const suggestedValueFor = (val, year, category) => {
    const y = parseInt(year);
    if (!isNaN(y) && y >= 2002 && category !== "") {
        try {
            return scoreHelp[year][category.toUpperCase()][val];
        } catch (e) {
            return "";
        }
    }
};

const suggestedFormat = year => {
    try {
        return formatHelp[year];
    } catch (e) {
        return "";
    }
};
