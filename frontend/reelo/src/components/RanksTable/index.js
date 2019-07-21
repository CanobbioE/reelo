import React from 'react';
import {
	Table,
	TableRow,
	TableBody,
	TableCell,
	TableHead,
	Paper,
	withStyles,
} from '@material-ui/core';

import './RanksTable.css';

const styles = theme => ({
	tableHeader: {
		backgroundColor: theme.palette.primary.main,
	},
	tableHeaderCell: {
		color: theme.palette.secondary.main,
	},
});

const RanksTable = props => {
	const {classes} = props;
	const renderHeader = () =>
		props.labels.map(label => (
			<TableCell className={classes.tableHeaderCell} key={label}>
				{label}
			</TableCell>
		));

	const renderRows = () =>
		props.rows.map((row, i) => {
			if (!isValidRow(row)) return null;
			return (
				<TableRow key={row.id} className={row.id % 2 ? 'ranks-table-row' : ''}>
					<TableCell> {i + 1}</TableCell>
					<TableCell>{row.name}</TableCell>
					<TableCell>{row.surname}</TableCell>
					<TableCell>{row.category}</TableCell>
					<TableCell>{row.reelo}</TableCell>
				</TableRow>
			);
		});

	return (
		<Paper className="scrollbar">
			<Table>
				<TableHead className={classes.tableHeader}>
					<TableRow>{renderHeader()}</TableRow>
				</TableHead>
				<TableBody>{renderRows()}</TableBody>
			</Table>
		</Paper>
	);
};

function isValidRow(row) {
	return (
		row.name !== '' &&
		row.surname !== '' &&
		row.cateogry !== '' &&
		row.reelo >= 0
	);
}

export default withStyles(styles)(RanksTable);
