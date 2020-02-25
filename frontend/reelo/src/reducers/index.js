import { combineReducers } from "redux";
import loginReducer from "./loginReducer";
import authReducer from "./authReducer";
import uploadReducer from "./uploadReducer";
import ranksReducer from "./ranksReducer";
import algorithmReducer from "./algorithmReducer";
import analysisReducer from "./analysisReducer";
import errrorReducer from "./errorReducer";
export default combineReducers({
    loginForm: loginReducer,
    auth: authReducer,
    uploadForm: uploadReducer,
    ranks: ranksReducer,
    algorithm: algorithmReducer,
    analysis: analysisReducer,
    errors: errrorReducer,
});
