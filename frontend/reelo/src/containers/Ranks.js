import { Grid, Typography, Fab, Button, InputBase } from "@material-ui/core";
import React, { useEffect, useState } from "react";
import RanksTable from "../components/RanksTable";
import DetailsTable from "../components/DetailsTable";
import ArrowForward from "@material-ui/icons/ArrowRight";
import { withStyles } from "@material-ui/core/styles";
import { connect } from "react-redux";
import { compose } from "redux";
import {
    fetchRanks,
    forceReelo,
    setRankPage,
    setRankSize,
    fetchTotalRanks,
    fetchAllYears,
} from "../actions";
import LoadingIcon from "../components/LoadingIcon";

const styles = theme => ({
    details: {
        color: "#f5f5f5",
        paddingLeft: "15px !important",
        marginLeft: "10px",
    },
    title: {
        marginTop: "28px",
        marginBottom: "15px",
    },
    inputRoot: {
        color: "inherit",
        border: "1px solid #f5f5f5",
    },
    inputInput: {
        padding: "1px",
        width: "100%",
    },
});

const Ranks = props => {
    const { classes } = props;
    const [details, setDetails] = useState(false);
    const [hovered, setHovered] = useState(-1);
    const [search, setSearch] = useState("");

    useEffect(() => {
        props.setRankPage(1);
        props.setRankSize(10);
        props.fetchTotalRanks();
        props.fetchAllYears();
        props.fetchRanks(1, 10);
    }, []);

    const handlePageChange = (event, page) => {
        props.setRankPage(page + 1);
        props.fetchRanks(page + 1, props.ranks.size);
    };

    const handleSizeChange = event => {
        props.setRankPage(1);
        const size = parseInt(event.target.value, 10);
        props.setRankSize(size);
        props.fetchRanks(1, size);
    };

    const handleSearch = event => {
        setSearch(event.target.value);
    };

    const rows = props.ranks.rows;
    const labels = ["#", "Nome", "Cognome", "Categoria", "Reelo"];
    var detailsLabels = [];
    if (props.ranks && props.ranks.years) {
        props.ranks.years.forEach(() => {
            detailsLabels = detailsLabels.concat([
                "Anno",
                "Categoria",
                "Esercizi",
                "Punteggio",
                //'Tempo',
                "Pre-REELO",
                "Posizione",
            ]);
        });
    }

    const detailsRows = populateDetails(rows, props.ranks.years, search);

    const ranksTable = (
        <Grid item container spacing={8} xs={details ? 4 : 10}>
            {!props.ranks.loading && (
                <Grid item xs={12}>
                    <RanksTable
                        filter={search}
                        onChangeRowsPerPage={handleSizeChange}
                        onChangePage={handlePageChange}
                        rows={rows}
                        labels={labels}
                        page={props.ranks.page}
                        count={props.ranks.count}
                        rowsPerPage={props.ranks.size}
                        onHover={setHovered}
                        hovered={hovered}
                    />
                </Grid>
            )}
        </Grid>
    );

    const detailsTable = (
        <Grid item xs={8}>
            {!props.ranks.loading && (
                <DetailsTable
                    filter={search}
                    onClose={() => setDetails(false)}
                    onHover={setHovered}
                    hovered={hovered}
                    rows={detailsRows}
                    labels={detailsLabels}
                />
            )}
        </Grid>
    );

    const detailsBtn = (
        <Grid item xs={2}>
            <Fab
                variant="extended"
                className={classes.details}
                onClick={() => setDetails(true)}
                size="small"
                color="primary">
                Dettagli
                <ArrowForward />
            </Fab>
        </Grid>
    );

    const searchbar = (
        <Grid item xs={12} className={classes.inputRoot}>
            <InputBase
                onChange={handleSearch}
                value={search}
                placeholder="Cerca..."
                fullWidth
                classes={{ input: classes.inputInput }}
            />
        </Grid>
    );

    const content = (
        <Grid
            container
            item
            spacing={8}
            xs={12}
            justify={details ? "flex-start" : "center"}
            alignItems={details ? "stretch" : "flex-start"}>
            {searchbar}
            {props.ranks.loading && (
                <Grid item xs={1}>
                    <LoadingIcon show={props.ranks.loading} />
                </Grid>
            )}
            {ranksTable}
            {!details ? detailsBtn : null}
            {details ? detailsTable : null}
        </Grid>
    );

    const error = (
        <Grid item xs={12}>
            <Typography align="center" variant="body2" color="error">
                Oops, si &egrave; verificato un errore:{" "}
                {rows ? props.errors.codeAsMessage : "Non sono presenti valori nella base di dati"}
            </Typography>
        </Grid>
    );

    return (
        <Grid container justify="center">
            <Grid item container spacing={24} xs={10}>
                <Grid item xs={12}>
                    <Typography variant="h4" className={classes.title}>
                        Classifiche
                    </Typography>
                </Grid>

                {props.errors.message === "" && rows ? content : error}
                <Grid item xs={12}>
                    {!props.auth.authenticated ? null : (
                        <Button
                            onClick={async () => {
                                await props.forceReelo();
                                await props.fetchRanks(props.ranks.page, 10);
                            }}
                            variant="contained"
                            color="primary">
                            Ricalcola REELO
                        </Button>
                    )}
                </Grid>
            </Grid>
        </Grid>
    );
};

const ndRow = (id, y) => ({
    id: `${id}-${y}`,
    year: y,
    category: "Non partecipato",
    e: "N/D",
    d: "N/D",
    // time: 0,
    pseudoReelo: "N/D",
    position: "N/D",
});

function populateDetails(rows, years, filter) {
    var detailsRows = [];
    if (rows) {
        rows.forEach(row => {
            if (!compliesToFilter(row, filter)) return null;
            var subRow = [];
            const h = row.history;
            years.forEach(y => {
                if (!h[y]) {
                    subRow = subRow.concat(ndRow(row.id, y));
                } else {
                    subRow = subRow.concat({
                        id: `${row.id}-${y}`,
                        year: y,
                        category: h[y].category.toUpperCase(),
                        e: `${h[y].e}/${h[y].eMax}=${(h[y].e / h[y].eMax).toFixed(2)}`,
                        d: `${h[y].d}/${h[y].dMax}=${(h[y].d / h[y].dMax).toFixed(2)}`,
                        //time: h[y].time > 0 ? h[y].time : 'N/D',
                        pseudoReelo: h[y].pseudoReelo.toFixed(0),
                        position: h[y].position,
                    });
                }
            });
            detailsRows.push(subRow);
        });
    }
    return detailsRows;
}

function compliesToFilter(row, filter) {
    return (
        row.name.toLowerCase().includes(filter.toLowerCase()) ||
        row.surname.toLowerCase().includes(filter.toLowerCase()) ||
        row.category.toLowerCase().includes(filter.toLowerCase())
    );
}

function mapStateToProps({ ranks, auth, errors }) {
    return { ranks, auth, errors };
}

const composedComponent = compose(
    connect(mapStateToProps, {
        fetchRanks,
        forceReelo,
        setRankPage,
        fetchTotalRanks,
        setRankSize,
        fetchAllYears,
    }),
);

export default withStyles(styles)(composedComponent(Ranks));
