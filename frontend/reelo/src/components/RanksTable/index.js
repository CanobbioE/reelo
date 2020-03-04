import React from "react";
import {
    Table,
    TableRow,
    TableBody,
    TableCell,
    TableHead,
    Paper,
    TableFooter,
    TablePagination,
} from "@material-ui/core";
import { withStyles } from "@material-ui/core/styles";
import "./RanksTable.css";

const styles = theme => ({
    tableHeader: {
        backgroundColor: theme.palette.primary.main,
    },
    tableHeaderCell: {
        color: theme.palette.secondary.main,
    },
    hover: {
        backgroundColor: "#d3bdfb",
    },
});

const RanksTable = props => {
    const { classes } = props;

    const renderHeader = () =>
        props.labels.map(label => (
            <TableCell className={classes.tableHeaderCell} key={label}>
                {label}
            </TableCell>
        ));

    const renderLabels = ({ from, to, count }) =>
        `Dal ${from} al ${to} di ${count} (pagina ${props.page} di ${(count / 10).toFixed()})`;

    const renderRows = () =>
        props.rows.map((row, i) => {
            if (!isValidRow(row)) return null;
            if (!compliesToFilter(row, props.filter)) return null;
            return (
                <TableRow
                    onMouseOver={() => props.onHover(row.id)}
                    key={row.id}
                    className={props.hovered === row.id ? classes.hover : i % 2 ? "grey-bg" : ""}>
                    <TableCell> {i + props.rowsPerPage * (props.page - 1) + 1}</TableCell>
                    <TableCell>{row.name}</TableCell>
                    <TableCell>{row.surname}</TableCell>
                    <TableCell>{row.category}</TableCell>
                    <TableCell>{row.reelo}</TableCell>
                </TableRow>
            );
        });

    return (
        <Paper className="paper scrollbar">
            <Table className="paper">
                <TableHead className={classes.tableHeader}>
                    <TableRow>{renderHeader()}</TableRow>
                </TableHead>
                <TableBody>{renderRows()}</TableBody>
                <TableFooter>
                    <TableRow>
                        <TablePagination
                            rowsPerPageOptions={[10, 50, 100, props.count]}
                            onChangeRowsPerPage={props.onChangeRowsPerPage}
                            rowsPerPage={props.rowsPerPage}
                            count={props.count}
                            page={props.page - 1}
                            onChangePage={props.onChangePage}
                            labelDisplayedRows={renderLabels}
                            labelRowsPerPage="Risultati per pagina:"
                        />
                    </TableRow>
                </TableFooter>
            </Table>
        </Paper>
    );
};

function compliesToFilter(row, filter) {
    return (
        row.name.toLowerCase().includes(filter.toLowerCase()) ||
        row.surname.toLowerCase().includes(filter.toLowerCase()) ||
        row.category.toLowerCase().includes(filter.toLowerCase())
    );
}

function isValidRow(row) {
    return row.name !== "" && row.surname !== "" && row.cateogry !== "" && row.reelo >= 0;
}

export default withStyles(styles)(RanksTable);
