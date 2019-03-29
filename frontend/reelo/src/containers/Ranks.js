import {Grid, Typography} from '@material-ui/core';
import React from 'react';
import {RanksTable} from '../components/RanksTable';

export default function Ranks(props) {
	const rows = [
		{id: 0, name: 'name', surname: 'surname', reelo: '123'},
		{id: 1, name: 'name', surname: 'surname', reelo: '123'},
		{id: 2, name: 'name', surname: 'surname', reelo: '123'},
	];

	const labels = ['Nome', 'Cognome', 'Reelo'];

	return (
		<Grid container justify="center">
			<Grid item xs={10}>
				<Typography variant="h4">Classifiche</Typography>
				<RanksTable rows={rows} labels={labels} />
			</Grid>
		</Grid>
	);
}
