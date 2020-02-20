import React from 'react';
import {Grid, TextField, Button, Typography} from '@material-ui/core';

export const EditAlgForm = props => {
	const handleChanges = fx => event => {
		fx(event.target.value);
	};

	const shouldDisable = props => {
		return (
			props.year === '' &&
			props.ex === '' &&
			props.final === '' &&
			props.mult === '' &&
			props.exp === '' &&
			props.np === ''
		);
	};

	return (
		<form onSubmit={props.onSubmit}>
			<Grid item container justify="center" spacing={24} xs={10}>
				<Grid item xs={10}>
					<TextField
						variant="outlined"
						fullWidth
						value={props.year}
						label="Anno di inizio"
						onChange={handleChanges(props.onYearChange)}
					/>
					<Typography variant="body1">
						<i>Valore attuale:</i> {props.currentValues.year}
					</Typography>
				</Grid>

				<Grid item xs={10}>
					<TextField
						variant="outlined"
						fullWidth
						value={props.ex}
						label="Coefficiente esercizi (K)"
						onChange={handleChanges(props.onExChange)}
					/>
					<Typography variant="body1">
						<i>Valore attuale:</i> {props.currentValues.ex}
					</Typography>
				</Grid>

				<Grid item xs={10}>
					<TextField
						variant="outlined"
						fullWidth
						value={props.final}
						label="Coefficiente finali internazionali (P)"
						onChange={handleChanges(props.onFinalChange)}
					/>
					<Typography variant="body1">
						<i>Valore attuale:</i> {props.currentValues.final}
					</Typography>
				</Grid>

				<Grid item xs={10}>
					<TextField
						variant="outlined"
						fullWidth
						value={props.mult}
						label="Fattore moltiplicativo (F)"
						onChange={handleChanges(props.onMultChange)}
					/>
					<Typography variant="body1">
						<i>Valore attuale:</i> {props.currentValues.mult}
					</Typography>
				</Grid>

				<Grid item xs={10}>
					<TextField
						variant="outlined"
						fullWidth
						value={props.exp}
						label="Anti exploit (AE)"
						onChange={handleChanges(props.onExpChange)}
					/>
					<Typography variant="body1">
						<i>Valore attuale: </i>
						{props.currentValues.exp}
					</Typography>
				</Grid>

				<Grid item xs={10}>
					<TextField
						variant="outlined"
						fullWidth
						value={props.np}
						label="Penalità di non-partecipazione"
						onChange={handleChanges(props.onNPChange)}
					/>
					<Typography variant="body1">
						<i>Valore attuale:</i> {props.currentValues.np}
					</Typography>
				</Grid>

				<Grid
					item
					container
					xs={12}
					justify="space-around"
					alignItems="flex-end">
					<Grid item xs={4}>
						<Button
							type="submit"
							variant="contained"
							color="primary"
							disabled={shouldDisable(props)}>
							Aggiorna variabili
						</Button>
					</Grid>

				</Grid>
			</Grid>
		</form>
	);
};
