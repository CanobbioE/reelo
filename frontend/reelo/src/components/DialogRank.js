import React from 'react';
import {
	Button,
	Dialog,
	DialogActions,
	DialogTitle,
	DialogContent,
	Table,
	TableHead,
	TableBody,
	TableCell,
	TableRow,
} from '@material-ui/core';

const DialogRank = props => {
	if (!props.data) return null;

	const history = props.data.history;
	const header = (
		<TableRow>
			<TableCell>Anno</TableCell>
			<TableCell>Categoria</TableCell>
			<TableCell>Esercizi (e)</TableCell>
			<TableCell>Mssimo numero di esercizi (eMax) </TableCell>
			<TableCell> Punteggio (d)</TableCell>
			<TableCell> Massimo punteggio (dMax) </TableCell>
			<TableCell> Pseudo REELO</TableCell>
		</TableRow>
	);
	const rows = Object.keys(history).map(k => (
		<TableRow key={k}>
			<TableCell>{k}</TableCell>
			<TableCell> {history[k].category} </TableCell>
			<TableCell> {history[k].e}</TableCell>
			<TableCell> {history[k].eMax}</TableCell>
			<TableCell> {history[k].d}</TableCell>
			<TableCell> {history[k].dMax}</TableCell>
			<TableCell> {history[k].pseudoReelo}</TableCell>
		</TableRow>
	));

	return (
		<Dialog open={props.open} fullScreen onClose={props.onClose}>
			<DialogTitle>
				{`${props.data.name} ${props.data.surname}   -   ${props.data.reelo}`}
			</DialogTitle>
			<DialogContent>
				<Table>
					<TableHead>{header}</TableHead>
					<TableBody>{rows}</TableBody>
				</Table>
			</DialogContent>

			<DialogActions>
				<Button
					onClick={props.onClose}
					color="primary"
					variant="contained"
					autoFocus>
					Chiudi
				</Button>
			</DialogActions>
		</Dialog>
	);
};

export default DialogRank;
