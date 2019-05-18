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
	List,
	ListItem,
	ListItemText,
} from '@material-ui/core';
import LoadingIcon from './LoadingIcon';

const categories = ['c1', 'c2', 'ce', 'cm', 'l1', 'l2', 'gp', 'hc'];

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

	const listPrimary = [
		'Nome',
		'Cognome',
		'Città',
		'Punti',
		'Tempo',
		'Esercizi',
	];
	const listSecondary = [
		'La colonna contenente il nome del concorrente (n)',
		'La colonna contenente il cognome del concorrente (c)',
		'La colonna in cui è specificata la sede in cui si sono tenuti i giochi (s)',
		'La colonna contenente il punteggio totale ottenuto dal concorrente (p)',
		'La colonna che specifica quanto tempo il concorrente ha impiegato a completare la prova',
		'La colonna che specifica il numero di esercizi risolti',
	];
	const fieldsList = listPrimary.map((text, i) => (
		<ListItem key={i}>
			<ListItemText
				primary={
					<React.Fragment>
						<Typography component="span" color="textPrimary" variant="body2">
							<strong>{text}</strong>
						</Typography>
					</React.Fragment>
				}
				secondary={
					<React.Fragment>
						<Typography component="span" color="textPrimary" variant="body2">
							{listSecondary[i]}
						</Typography>
					</React.Fragment>
				}
			/>
		</ListItem>
	));
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
						Inserisci il tipo di dati contenuto nelle colonne in modo ordinato e
						separato da spazi, i possibili valori sono:
					</Typography>
					<List>{fieldsList}</List>
					<Typography variant="body2">
						I valori tra le parentesi sono abbreviazioni che possono essere
						inserite invece dell'intera parola. Specifica solo le colonne che
						compaiono, ad esempio se ci fosse solo il nome scrivi solo il valore
						"nome" (senza virgolette)
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

				<Grid item container xs={12} justify="flex-start" alignItems="flex-end">
					<Grid item xs={2}>
						<Button type="submit" variant="contained" color="primary">
							Carica
						</Button>
					</Grid>
					<Grid item xs={1}>
						<LoadingIcon show={props.loading} />
					</Grid>
				</Grid>
			</Grid>
		</form>
	);
};
