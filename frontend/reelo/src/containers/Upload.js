import React, {useState} from 'react';
import {connect} from 'react-redux';
import {compose} from 'redux';
import {
	Grid,
	Typography,
	DialogContentText,
	DialogContent,
} from '@material-ui/core';
import {UploadForm} from '../components/UploadForm';
import RequireAuth from './RequireAuth';
import DialogAlert from '../components/DialogAlert';
import {
	updateUploadFile,
	updateUploadYear,
	updateUploadCategory,
	updateUploadFormat,
	updateUploadStart,
	updateUploadEnd,
	updateUploadIsParis,
	uploadFile,
	resetUploadForm,
} from '../actions';

function Upload(props) {
	const [alertOpen, setAlertOpen] = useState(props.uploadForm.error !== '');
	const info = (
		<Grid item xs={10}>
			<Typography variant="subtitle1">
				Utilizza questa pagina per inserire un file classifica con estensione
				".txt".
			</Typography>
			<Typography variant="subtitle2">
				Utilizza l'apposita sezione per specificare la disposizione dei dati
				all'interno del file. (e.g. "nome cognome sede punteggio tempo").
			</Typography>
			<br />
		</Grid>
	);

	const error = props.uploadForm.error.split('\n').map((item, i) => {
		return (
			<span key={i}>
				{item}
				<br />
			</span>
		);
	});

	const dialogContent = (
		<DialogContent>
			<DialogContentText id="alert-dialog-description">
				Durante la lettura del file si sono verificati degli errori, ecco il
				messaggio generato:
			</DialogContentText>
			<DialogContentText color="error" id="alert-dialog-description">
				{error}
			</DialogContentText>
			<DialogContentText id="alert-dialog-description">
				Se possibile cerca di sistemare il documento.
			</DialogContentText>
		</DialogContent>
	);

	if (props.uploadForm.error !== '' && !alertOpen) setAlertOpen(true);

	return (
		<Grid container justify="center">
			{info}
			<Grid container item xs={10}>
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
					onStartInput={props.updateUploadStart}
					startValue={props.uploadForm.start}
					onEndInput={props.updateUploadEnd}
					endValue={props.uploadForm.end}
					loading={props.uploadForm.loading}
					startSugg={props.uploadForm.startSugg}
					endSugg={props.uploadForm.endSugg}
					formatSugg={props.uploadForm.formatSugg}
				/>
				<DialogAlert
					open={alertOpen}
					onClose={() => {
						props.resetUploadForm();
						setAlertOpen(false);
					}}
					content={dialogContent}
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
			resetUploadForm,
			updateUploadStart,
			updateUploadEnd,
		},
	),
);

export default composedComponent(Upload);
