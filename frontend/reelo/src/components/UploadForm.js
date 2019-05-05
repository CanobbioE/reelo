import React from 'react';
import {Grid, Input, TextField, Button} from '@material-ui/core';

export const UploadForm = props => {
	const handleSubmit = event => {
		event.preventDefault();
		props.onSubmit();
	};

	return (
		<form onSubmit={handleSubmit}>
			<Grid item container spacing={24}>
				<Grid item xs={12}>
					<Input type="file" disableUnderline />
				</Grid>
				<Grid item xs={12}>
					<TextField label="Formato dati" />
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
