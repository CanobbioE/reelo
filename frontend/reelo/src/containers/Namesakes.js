import React, { useState } from "react";
import { connect } from "react-redux";
import { compose } from "redux";
import { Grid, Typography, Button } from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import { fetchNamesakes, updateNamesake, acceptNamesake, commentNamesake } from "../actions";
import NamesakeForm from "../components/NamesakeForm";
import LoadingIcon from "../components/LoadingIcon";

const styles = theme => ({
    title: {
        marginTop: "28px",
        marginBottom: "15px",
        paddingLeft: 5 * theme.spacing.unit,
    },
    subtitle: {
        marginTop: "28px",
        marginBottom: "15px",
    },
});

const Namesakes = props => {
    const { classes } = props;

    const [selection, setSelection] = useState([]);
    const toggleSelection = index => (namesake, isSelected) => {
        if (isSelected) {
            var tmp = [];
            if (selection && selection.length > 0) {
                tmp = selection.slice(0, index).concat(selection.slice(index + 1));
            } else {
                tmp = [namesake];
            }
            setSelection(tmp);
        } else {
            const tmp = selection.slice();
            tmp.push(namesake);
            setSelection(tmp);
        }
    };

    const handleMerge = i => merge => {
        setSelection([])
        const ret = merge(props.analysis.namesakes[i + 1]);
        if (ret) {
            props.updateNamesake(i, ret);
        }
    };

    const handleAccept = async () => {
        if (!window.confirm("Sicuro di voler confermare la soluzione proposta?")) {
            return;
        }
        selection.forEach(async namesake => await props.acceptNamesake(namesake));
        props.fetchNamesakes(1, -1);
        setSelection([])
    };

    const handleComment = async (namesake, comment) => {
        await props.commentNamesake(namesake, comment);
        props.fetchNamesakes(1, -1);
    };

    const renderNamesakes = () =>
        props.analysis &&
        props.analysis.namesakes &&
        props.analysis.namesakes.map((namesake, i) => (
            <NamesakeForm
                key={`${namesake.player.id} ${namesake.id}`}
                namesake={namesake}
                onComment={handleComment}
                onMerge={handleMerge(i)}
                onSelect={toggleSelection(i)}
            />
        ));

    const renderError = () => (
        <Grid item xs={12}>
            <Typography color="error" align="center">
                {props.errors.codeAsMessage} <br />
                {props.errors.code === "E_NO_AUTH" || props.errors.code === "E_BAD_REQ"
                    ? props.errors.message
                    : null}
            </Typography>
        </Grid>
    );

    const infoNamesakes = num =>
        props.analysis.loading
            ? "Calcolo omonimi in corso..."
            : num <= 0
            ? "Non ho trovato omonimi, hai premuto il pulsante qui sopra?"
            : `Ho trovato ${num} omonim${num > 1 ? "i" : "o"}`;

    return (
        <Grid container justify="center">
            <Grid item xs={11}>
                <Typography variant="h4" className={classes.title}>
                    Risoluzione Omonimi
                </Typography>
            </Grid>
            <Grid item xs={11}>
                <Typography variant="subtitle1">
                    Utilizza questa pagina per risolvere i casi di omonimia. Ci sono quattro azioni
                    disponibili:
                    <dl>
                        <dt>
                            <b>Accetta Selezionati:</b>
                        </dt>
                        <dd>
                            Conferma la soluzione proposta e crea un nuovo giocatore con la
                            cronologia di partecipazioni proposta per ogni soluzione selezionata.
                        </dd>

                        <dt>
                            <b>Seleziona</b>
                        </dt>
                        <dd>
                            Aggiunge la soluzione proposta all'elenco di soluzioni accettabili. Una
                            volta soddisfatti con la propria selezione, si pu&ograve; procedere
                            premendo il pulsante "Accetta Selezionati".
                        </dd>

                        <dt>
                            <b>Unisci:</b>
                        </dt>
                        <dd>
                            La cronologia del giocatore selezionato viene unita alla cronologia del
                            suo omonimo sottostante. Se i due giocatori non sono omonimi non succede
                            nulla. La modifica viene attuata solo una volta premuto il pulsante
                            "Accetta".
                        </dd>

                        <dt>
                            <b>Commenta:</b>
                        </dt>
                        <dd>
                            Semplicemente aggiunge un commento in modo da poi risolvere manualmente
                            l'omonimia. Nessuna modifica drastica viene applicata. Non necessita che
                            venga premuto il tasto "Accetta".
                        </dd>
                    </dl>
                </Typography>
                <hr />
                <Grid item container spacing={8} justify="center" alignItems="center">
                    <Grid item>Elenca tutti gli omonimi non risolvibili automaticamente:</Grid>
                    <Grid item>
                        <Button
                            variant="contained"
                            size="small"
                            disabled={props.analysis.loading}
                            onClick={() => props.fetchNamesakes(1, -1)}>
                            Mostra tutto
                        </Button>
                    </Grid>
                    <Grid item>
                        <Button
                            variant="contained"
                            color="primary"
                            size="small"
                            disabled={selection.length <= 0}
                            onClick={handleAccept}>
                            Accetta selezionati ({selection.length})
                        </Button>
                    </Grid>
                    <Grid item xs={12}>
                        {props.analysis && props.analysis.namesakes && (
                            <Typography align="center">
                                {infoNamesakes(props.analysis.namesakes.length)}
                            </Typography>
                        )}
                    </Grid>
                </Grid>
            </Grid>

            <Grid item container justify="center" spacing={24} xs={10}>
                <Grid item xs={1}>
                    <Typography className={classes.subtitle} align="left" variant="h6">
                        Giocatore
                    </Typography>
                </Grid>
                <Grid item xs={6}>
                    <Typography className={classes.subtitle} align="center" variant="h6">
                        Soluzione proposta
                    </Typography>
                </Grid>
                <Grid item xs={2}>
                    <Typography className={classes.subtitle} align="left" variant="h6">
                        Commento
                    </Typography>
                </Grid>
                <Grid item xs={3}>
                    <Typography className={classes.subtitle} align="center" variant="h6">
                        Azioni
                    </Typography>
                </Grid>
                {props.errors.message !== "" ? renderError() : null}
                <Grid item xs={1}>
                    <LoadingIcon show={props.analysis.loading} />
                </Grid>
                {props.analysis.loading && props.errors.message === "" ? null : renderNamesakes()}
            </Grid>
        </Grid>
    );
};

function mapStateToProps({ uploadForm, analysis, errors }) {
    return { uploadForm, analysis, errors };
}

const composedComponent = compose(
    // TODO RequireAuth,
    connect(mapStateToProps, {
        fetchNamesakes,
        updateNamesake,
        acceptNamesake,
        commentNamesake,
    }),
);

export default withStyles(styles)(composedComponent(Namesakes));
