import React from "react";
import {
    Table,
    TableRow,
    TableBody,
    TableCell,
    TableHead,
    Paper,
    TableFooter,
} from "@material-ui/core";
import "./RanksTable/RanksTable.css";
import { withStyles } from "@material-ui/core/styles";

const styles = theme => ({
    tableHeader: {
        backgroundColor: theme.palette.primary.main,
        cursor: "pointer",
    },
    tableHeaderCell: {
        color: theme.palette.secondary.main,
    },
    divider: {
        borderRight: "2px solid black",
    },
    dividerHead: {
        color: theme.palette.secondary.main,
        borderRight: "2px solid black",
    },
    greyBg: {
        backgroundColor: "#f5f5f5",
    },
    hover: {
        backgroundColor: "#d3bdfb",
    },
});

const DetailsTable = props => {
    const { classes } = props;

    const renderHeader = () =>
        props.labels.map((label, i) => (
            <TableCell
                key={i}
                className={label === "Posizione" ? classes.dividerHead : classes.tableHeaderCell}>
                {label}
            </TableCell>
        ));

    const renderRows = () =>
        props.rows.map((row, index) => {
            return (
                <TableRow
                    key={`${row.id}-${index}`}
                    onMouseOver={() => props.onHover(parseInt(row[0].id.split("-")[0], 10))}
                    className={
                         props.hovered === parseInt(row[0].id.split("-")[0], 10)
                            ? classes.hover
                            : index % 2
                            ? classes.greyBg
                            : ""
                    }>
                    {row.map((subRow, i) =>
                        Object.keys(subRow).map((k, j) =>
                            k === "id" ? null : (
                                <TableCell
                                    key={`${row.id}-${i}-${j}`}
                                    className={k === "position" ? classes.divider : ""}>
                                    {subRow[k]}
                                </TableCell>
                            ),
                        ),
                    )}
                </TableRow>
            );
        });

    return (
        <Paper className="paper scrollbar">
            <Table className="paper">
                <TableHead className={classes.tableHeader} onClick={props.onClose}>
                    <TableRow>{renderHeader()}</TableRow>
                </TableHead>
                <TableBody>{renderRows()}</TableBody>
                <TableFooter>
                    <TableRow></TableRow>
                </TableFooter>
            </Table>
        </Paper>
    );
};

export default withStyles(styles)(DetailsTable);
