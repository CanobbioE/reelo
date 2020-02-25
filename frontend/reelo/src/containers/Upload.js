import React, { useState } from "react";
import { withStyles } from "@material-ui/core/styles";
import { connect } from "react-redux";
import { compose } from "redux";
import { Grid, Typography, DialogContentText, DialogContent } from "@material-ui/core";
import UploadForm from "../components/UploadForm";
import RequireAuth from "./RequireAuth";
import DialogAlert from "../components/DialogAlert";
import {
    updateUploadFile,
    updateUploadYear,
    updateUploadCategory,
    updateUploadFormat,
    updateUploadStart,
    updateUploadEnd,
    updateUploadIsParis,
    uploadFile,
    resetUploadForm,
    checkExistence,
} from "../actions";

const styles = () => ({
    title: {
        marginTop: "28px",
        marginBottom: "15px",
    },
});

function Upload(props) {
    const { classes } = props;
    const [alertOpen, setAlertOpen] = useState(props.errors.message !== "");
    // const [dialogOpen, setDialogOpen] = useState(false);
    const handleSubmit = async (file, cat, y, isParis, fmt, s, e) => {
        var exists = await props.checkExistence(y, cat, isParis);
        if (exists) {
            if (window.confirm("stai caricando una classifica già presente")) {
                props.uploadFile(file, cat, y, isParis, fmt, s, e);
            }
        } else {
            props.uploadFile(file, cat, y, isParis, fmt, s, e);
        }
    };
    const info = (
        <Grid item xs={10}>
            <Typography variant="h4" className={classes.title}>
                Caricamento
            </Typography>
            <Typography variant="subtitle1">
                Utilizza questa pagina per inserire un file classifica con estensione ".txt".
            </Typography>
            <br />
        </Grid>
    );

    const error = props.errors.message.split("\n").map((item, i) => {
        return (
            <span key={i}>
                {item}
                <br />
            </span>
        );
    });

    const dialogContent = (
        <DialogContent>
            <DialogContentText id="alert-dialog-description">
                Durante la lettura del documento si sono verificati degli errori, ecco il messaggio
                generato:
            </DialogContentText>
            <DialogContentText color="error" id="alert-dialog-description">
                {error}
            </DialogContentText>
            <DialogContentText id="alert-dialog-description">
                Se possibile cerca di sistemare il documento. Ricordando che a volte i nomi/cognomi
                multipli non vengono riconosciuti. Per ovviare al problema sostituisci gli spazi tra
                i nomi/cognomi multipli con dei trattini bassi. <br />
                Ad esempio "maria giovanna da vinci" diventerà "maria_giovanna da_vinci"
            </DialogContentText>
        </DialogContent>
    );

    if (props.errors.message !== "" && !alertOpen) setAlertOpen(true);

    return (
        <Grid container justify="center">
            {info}
            <Grid container item xs={10}>
                <UploadForm
                    onSubmit={handleSubmit}
                    onFileInput={props.updateUploadFile}
                    fileValue={props.uploadForm.file}
                    onFormatInput={props.updateUploadFormat}
                    formatValue={props.uploadForm.format}
                    onYearInput={props.updateUploadYear}
                    yearValue={props.uploadForm.year}
                    onCategoryInput={props.updateUploadCategory}
                    categoryValue={props.uploadForm.category}
                    onIsParisInput={props.updateUploadIsParis}
                    isParisValue={props.uploadForm.isParis}
                    onStartInput={props.updateUploadStart}
                    startValue={props.uploadForm.start}
                    onEndInput={props.updateUploadEnd}
                    endValue={props.uploadForm.end}
                    loading={props.uploadForm.loading}
                    startSugg={props.uploadForm.startSugg}
                    endSugg={props.uploadForm.endSugg}
                    formatSugg={props.uploadForm.formatSugg}
                />
                <DialogAlert
                    open={alertOpen}
                    onClose={() => {
                        props.resetUploadForm();
                        setAlertOpen(false);
                    }}
                    content={dialogContent}
                />
            </Grid>
        </Grid>
    );
}

function mapStateToProps({ uploadForm, errors }) {
    return { uploadForm, errors };
}

const composedComponent = compose(
    RequireAuth,
    connect(mapStateToProps, {
        updateUploadFile,
        updateUploadYear,
        updateUploadCategory,
        uploadFile,
        updateUploadFormat,
        updateUploadIsParis,
        resetUploadForm,
        updateUploadStart,
        updateUploadEnd,
        checkExistence,
    }),
);

export default withStyles(styles)(composedComponent(Upload));
