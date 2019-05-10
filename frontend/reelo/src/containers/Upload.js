import React from 'react';
import {connect} from 'react-redux';
import {compose} from 'redux';
import {Grid, Typography} from '@material-ui/core';
import {UploadForm} from '../components/UploadForm';
import RequireAuth from './RequireAuth';
import {
	updateUploadFile,
	updateUploadYear,
	updateUploadCategory,
	updateUploadFormat,
	updateUploadIsParis,
	uploadFile,
} from '../actions';

function Upload(props) {
	const info = (
		<Grid item xs={10}>
			<Typography variant="body2">
				Utilizza questa pagina per inserire un file classifica con estensione
				".txt".
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
				<UploadForm
					onSubmit={props.uploadFile}
					onFileInput={props.updateUploadFile}
					fileValue={props.uploadForm.file}
					onFormatInput={props.updateUploadFormat}
					formatValue={props.uploadForm.format}
					onYearInput={props.updateUploadYear}
					yearValue={props.uploadForm.year}
					onCategoryInput={props.updateUploadCategory}
					categoryValue={props.uploadForm.category}
					onIsParisInput={props.updateUploadIsParis}
					isParisValue={props.uploadForm.isParis}
				/>
			</Grid>
		</Grid>
	);
}

function mapStateToProps({uploadForm}) {
	return {uploadForm};
}

const composedComponent = compose(
	RequireAuth,
	connect(
		mapStateToProps,
		{
			updateUploadFile,
			updateUploadYear,
			updateUploadCategory,
			uploadFile,
			updateUploadFormat,
			updateUploadIsParis,
		},
	),
);

export default composedComponent(Upload);
