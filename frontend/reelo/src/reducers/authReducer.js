import { AUTH_USER } from "../utils/Types";

const INITIAL_STATE = {
    authenticated: localStorage.getItem("token") !== "" && localStorage.getItem("token"),
};

export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case AUTH_USER:
            return {
                ...state,
                authenticated: action.payload,
            };
        default:
            return state;
    }
};
