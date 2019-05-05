import {Grid, Typography} from '@material-ui/core';
import React from 'react';

export default function Home(props) {
	return (
		<Grid container justify="center">
			<Grid item xs={10}>
				<Typography variant="h4">Reelo</Typography>
				<Typography variant="subtitle1">
					<em>
						Il nome REELO vuole richiamare il sistema ELO e in Esperanto
						significa "numero reale".
					</em>
				</Typography>

				<dl>
					<dt>
						<strong> Ideazione:</strong>
					</dt>
					<dd> Cesco Reale </dd>
					<dt>
						<strong>Implementazione: </strong>
					</dt>
					<dd> Fabio Angelini </dd>
					<dd> Anna Bernardi </dd>
					<dd> Edoardo Canobbio </dd>
					<dt>
						<strong>Comitato scientifico: </strong>
					</dt>
					<dd>
						Fabio Angelini, Marco Broglia, Andrea Nari, Marco Pellegrini, Cesco
						Reale
					</dd>
					<dt>
						<strong>Collaboratori: </strong>
					</dt>
					<dd>
						David Barbato, Mirko Cappuccia, Maurizio De Leo, Francesco Moradin,
						Alberto Saracco
					</dd>
				</dl>
			</Grid>
		</Grid>
	);
}
