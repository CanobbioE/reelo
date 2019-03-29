import {
	Grid,
	Typography,
	Table,
	TableRow,
	TableBody,
	TableCell,
	TableHead,
} from '@material-ui/core';
import React from 'react';

export default function Ranks(props) {
	const rows = [
		{id: 0, name: 'name', surname: 'surname', reelo: '123'},
		{id: 1, name: 'name', surname: 'surname', reelo: '123'},
		{id: 2, name: 'name', surname: 'surname', reelo: '123'},
	];

	const labels = ['Nome', 'Cognome', 'Reelo'];

	const renderHeader = () =>
		labels.map(label => <TableCell> {label} </TableCell>);

	const renderRows = () =>
		rows.map(row => (
			<TableRow key={row.id}>
				<TableCell>{row.name}</TableCell>
				<TableCell>{row.surname}</TableCell>
				<TableCell>{row.reelo}</TableCell>
			</TableRow>
		));

	return (
		<Grid container justify="center">
			<Grid item xs={10}>
				<Typography variant="h4">Classifiche</Typography>
				<Table>
					<TableHead>
						<TableRow>{renderHeader()}</TableRow>
					</TableHead>
					<TableBody>{renderRows()}</TableBody>
				</Table>
			</Grid>
		</Grid>
	);
}
