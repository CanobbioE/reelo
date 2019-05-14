import React from 'react';
import {
	Grid,
	Input,
	Select,
	MenuItem,
	TextField,
	Button,
	Typography,
	Checkbox,
} from '@material-ui/core';

const categories = ['c1', 'c2', 'l1', 'l2', 'gp', 'hc'];

export const UploadForm = props => {
	const handleSubmit = event => {
		event.preventDefault();
		props.onSubmit(
			props.fileValue,
			props.categoryValue,
			props.yearValue,
			props.isParisValue,
			props.formatValue,
		);
	};

	const handleChanges = fx => event => {
		fx(event.target.value);
	};
	const handleCheckbox = fx => event => {
		fx(event.target.checked);
	};
	const handleFileSelection = fx => event => {
		fx(event.target.files[0]);
	};

	const selectItems = [
		<MenuItem key="-1" value="-1">
			<em>Seleziona una categoria</em>
		</MenuItem>,
	];
	selectItems.push(
		categories.map(category => (
			<MenuItem key={category} value={category}>
				{category.toUpperCase()}
			</MenuItem>
		)),
	);

	return (
		<form onSubmit={handleSubmit}>
			<Grid item container spacing={24}>
				<Grid item xs={12}>
					<Input
						label="Documento"
						type="file"
						onChange={handleFileSelection(props.onFileInput)}
						disableUnderline
						required
					/>
				</Grid>

				<Grid item xs={12}>
					<Typography variant="body2">
						Inserisci il tipo di dati nelle colonne in modo ordinato e separato
						da spazi: e.g. nome cognome città esercizi punteggio tempo
					</Typography>
					<TextField
						required
						value={props.formatValue}
						label="Formato dati"
						onChange={handleChanges(props.onFormatInput)}
					/>
				</Grid>

				<Grid item xs={12}>
					<Typography variant="body2">
						Inserisci l'anno a cui i risultati fanno riferimento
					</Typography>
					<TextField
						required
						value={props.yearValue}
						label="Anno"
						onChange={handleChanges(props.onYearInput)}
					/>
				</Grid>

				<Grid item xs={12}>
					<Typography variant="body2">
						Seleziona la categoria a cui il documento fa riferimento
					</Typography>
					<Select
						required
						value={props.categoryValue || '-1'}
						onChange={handleChanges(props.onCategoryInput)}
						inputProps={{
							name: 'Categoria',
							id: 'category',
						}}>
						{selectItems}
					</Select>
				</Grid>

				<Grid item container xs={12} alignItems="baseline">
					<Grid item xs={1}>
						<Checkbox
							checked={props.isParisValue}
							onChange={handleCheckbox(props.onIsParisInput)}
							value={props.isParisValue + ''}
							color="primary"
						/>
					</Grid>
					<Grid item xs={11}>
						<Typography variant="body2">
							Seleziona questa casella se la classifica fa riferimento a
							risultati internazionali (Parigi).
						</Typography>
					</Grid>
				</Grid>

				<Grid item xs={12}>
					<Button type="submit" variant="outlined">
						Carica
					</Button>
				</Grid>
			</Grid>
		</form>
	);
};
