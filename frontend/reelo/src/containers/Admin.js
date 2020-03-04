import { Grid } from "@material-ui/core";
import React, { useEffect } from "react";
import LoginForm from "../components/LoginForm";
import Logout from "../components/Logout";
import { updateEmail, updatePassword, signin, signout } from "../actions";
import { connect } from "react-redux";
import Globals from "../config/Globals";

const Admin = props => {
    useEffect(() => {
        if (props.auth.authenticated) {
            props.history.push(Globals.routes.home);
        }
    }, [props.auth.authenticated]);
    const login = async event => {
        event.preventDefault();
        await props.signin(props.loginForm.email, props.loginForm.password);
        if (props.auth.authenticated) {
            props.history.push(Globals.routes.home);
        }
    };

    const logout = event => {
        event.preventDefault();
        props.signout();
        props.history.push(Globals.routes.home);
    };

    const error =
        props.errors.message === "" ? null : (
            <>
                {props.errors.codeAsMessage}
                <br />
                messaggio dal server: {props.errors.message}
            </>
        );
    const loginForm = (
        <LoginForm
            error={error}
            onPasswordChange={props.updatePassword}
            onEmailChange={props.updateEmail}
            onSubmit={login}
            emailValue={props.loginForm.email}
            passwordValue={props.loginForm.password}
        />
    );
    const logoutForm = <Logout onClick={logout} />;
    const form = props.auth.authenticated ? logoutForm : loginForm;

    return (
        <Grid container justify="center">
            {form}
        </Grid>
    );
};

function mapStateToProps({ loginForm, auth, errors }) {
    return { loginForm, auth, errors };
}

const composedComponent = connect(mapStateToProps, {
    updateEmail,
    updatePassword,
    signin,
    signout,
});

export default composedComponent(Admin);
