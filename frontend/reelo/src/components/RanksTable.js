import React from 'react';
import {
	Table,
	TableRow,
	TableBody,
	TableCell,
	TableHead,
} from '@material-ui/core';

export const RanksTable = props => {
	const renderHeader = () =>
		props.labels.map(label => <TableCell key={label}> {label} </TableCell>);

	const renderRows = () =>
		props.rows.map(row => (
			<TableRow key={row.id}>
				<TableCell>{row.name}</TableCell>
				<TableCell>{row.surname}</TableCell>
				<TableCell>{row.reelo}</TableCell>
			</TableRow>
		));

	return (
		<Table>
			<TableHead>
				<TableRow>{renderHeader()}</TableRow>
			</TableHead>
			<TableBody>{renderRows()}</TableBody>
		</Table>
	);
};
