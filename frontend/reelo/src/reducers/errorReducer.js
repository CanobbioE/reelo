import { ERROR, ERROR_RESET } from "../utils/Types";

const INITIAL_STATE = {
    message: "",
    code: "",
    codeAsMessage: "",
    httpStatus: 0,
};

export default (state = INITIAL_STATE, action) => {
    switch (action.type) {
        case ERROR:
            return { ...state, ...parseError(action.payload) };
        case ERROR_RESET:
            return { ...INITIAL_STATE };
        default:
            return state;
    }
};

const parseError = ({ message, code, status }) => {
    const err = {
        code: `${code}`,
        message: `${message}`,
        status: status,
        codeAsMessage: "",
    };
    switch (code) {
        case "E_DB_UPDATE":
            return { ...err, codeAsMessage: "impossibile aggiornare la base di dati" };
        case "E_DB_CREATE":
            return { ...err, codeAsMessage: "impossibile inserire valori nella base di dati" };
        case "E_DB_FIND":
            return { ...err, codeAsMessage: "impossibile trovare valori nella base di dati" };
        case "E_GENERIC":
            return { ...err, codeAsMessage: "è avvenuto un errore inaspettato" };
        case "E_NO_AUTH":
            return { ...err, codeAsMessage: "credenziali non valide" };
        case "E_BAD_REQ":
            return { ...err, codeAsMessage: "richiesta non valida" };
        case "E_PARSE_WARN":
            return { ...err, codeAsMessage: "impossibile analizzare il documento" };
        default:
            return { ...err, codeAsMessage: "errore inaspettato" };
    }
};
