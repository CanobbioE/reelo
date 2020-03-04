import { NAMESAKE_FETCH_LOADING, NAMESAKE_FETCH_SUCCESS, NAMESAKE_UPDATE } from "../utils/Types";

const INITIAL_STATE = {
    loading: false,
    namesakes: [],
    fixedNamesakes: {},
};

export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case NAMESAKE_FETCH_LOADING:
            return { ...state, namesakes: [], loading: true };
        case NAMESAKE_FETCH_SUCCESS:
            return { ...state, namesakes: action.payload, loading: false };
        case NAMESAKE_UPDATE:
            return {
                ...state,
                namesakes: updateSolvers(
                    action.payload.index,
                    action.payload.namesake,
                    state.namesakes,
                ),
            };
        default:
            return { ...state };
    }
};

const updateSolvers = (index, namesake, namesakes) => {
    let tmp = namesakes.slice(0, index + 1);
    tmp[index] = namesake;
    tmp = tmp.concat(namesakes.slice(index + 2));
    return tmp;
};
