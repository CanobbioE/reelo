import React from 'react';
import {Grid, Typography} from '@material-ui/core';
import {UploadForm} from '../components/UploadForm';

export default function Upload(props) {
	let info = (
		<Grid item xs={10}>
			<Typography variant="body2">
				Utilizza questa pagina per inserire un file classifica con estensione
				".txt" e nome in formato "anno_categoria".
			</Typography>
			<Typography variant="body2">
				Le categorie possibili sono: C1, C2, L1, L2, GP.
			</Typography>
			<Typography variant="body2">
				Utilizza l'apposita sezione per specificare la disposizione dei dati
				all'interno del file. (e.g. "nome cognome sede punteggio tempo").
			</Typography>
			<br />
		</Grid>
	);

	return (
		<Grid container justify="center">
			{info}
			<Grid item xs={10}>
				<UploadForm onSubmit={props.uploadFile} />
			</Grid>
		</Grid>
	);
}
