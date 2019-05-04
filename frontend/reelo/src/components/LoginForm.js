import {Grid, TextField, Button} from '@material-ui/core';
import React from 'react';

export default function LoginForm(props) {
	return (
		<form onSubmit={props.onSubmit}>
			<Grid item xs={12}>
				<TextField label="Posta elettronica" />
			</Grid>
			<Grid item xs={12}>
				<TextField label="Parola chiave" />
			</Grid>

			<Grid item xs={12}>
				<Button type="submit" variant="outlined">
					Accedi
				</Button>
			</Grid>
		</form>
	);
}
