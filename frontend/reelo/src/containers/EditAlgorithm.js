import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import LoadingIcon from "../components/LoadingIcon";
import { compose } from "redux";
import { Grid, Typography, Dialog, DialogActions, DialogContent, Button } from "@material-ui/core";
import RequireAuth from "./RequireAuth";
import { EditAlgForm } from "../components/EditAlgForm";
import Globals from "../config/Globals";
// import doneImg from '../assets/images/done-checkmark.png';
import { withStyles } from "@material-ui/core/styles";
import {
    updateAlgYear,
    updateAlgEx,
    updateAlgFinal,
    updateAlgMult,
    updateAlgExp,
    updateAlgNP,
    fetchVars,
    updateAlg,
} from "../actions";

const styles = () => ({
    title: {
        marginTop: "28px",
        marginBottom: "15px",
    },
});

function EditAlgorithm(props) {
    const { classes } = props;
	const [done, setDone] = useState(false);
    useEffect(() => {
        props.fetchVars();
    }, []);

    const handleSubmit = async event => {
        event.preventDefault();
        await props.updateAlg(
            props.algorithm.year,
            props.algorithm.ex,
            props.algorithm.final,
            props.algorithm.mult,
            props.algorithm.exp,
            props.algorithm.np,
            props.algorithm.currentValues,
        );
        if (props.errors.message === "") {
            setDone(true);
        }
        props.history.push(Globals.routes.varchange);
        window.scrollTo(0, 0);
    };

    const info = (
        <Grid item xs={10}>
            <Typography variant="h4" className={classes.title}>
                Modifica algoritmo
            </Typography>

            <Typography variant="subtitle1">
                Utilizza questa pagina per modificare alcune delle costanti che vengono utilizzate
                nell'algoritmo per il calcolo del REELO
            </Typography>
            <br />
        </Grid>
    );

	const handleClose = () => {
		setDone(false)
		props.fetchVars()
	}
    return (
        <Grid container justify="center">
            {info}
            <Grid item container direction="column" alignItems="center">
                <Dialog open={done} keepMounted onClose={handleClose}>
                    <DialogContent>
                        <LoadingIcon show={props.loading} />
                        {props.errors.message !== "" ? (
                            <Typography color="error"> {props.errors.codeAsMessage}</Typography>
                        ) : (
                            <Typography>Modifica avvenuta con successo</Typography>
                        )}
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose} color="primary">
                            Ok
                        </Button>
                    </DialogActions>
                </Dialog>
            </Grid>
            <Grid item container xs={10}>
                <EditAlgForm
                    onYearChange={props.updateAlgYear}
                    onExChange={props.updateAlgEx}
                    onFinalChange={props.updateAlgFinal}
                    onMultChange={props.updateAlgMult}
                    onExpChange={props.updateAlgExp}
                    onNPChange={props.updateAlgNP}
                    onSubmit={handleSubmit}
                    year={props.algorithm.year}
                    ex={props.algorithm.ex}
                    final={props.algorithm.final}
                    mult={props.algorithm.mult}
                    exp={props.algorithm.exp}
                    np={props.algorithm.np}
                    loading={props.algorithm.loading}
                    currentValues={props.algorithm.currentValues}
                />
            </Grid>
        </Grid>
    );
}
function mapStateToProps({ algorithm, errors }) {
    return { algorithm, errors };
}

const composedComponent = compose(
    RequireAuth,
    connect(mapStateToProps, {
        updateAlgYear,
        updateAlgEx,
        updateAlgFinal,
        updateAlgMult,
        updateAlgExp,
        updateAlgNP,
        fetchVars,
        updateAlg,
    }),
);

export default withStyles(styles)(composedComponent(EditAlgorithm));
