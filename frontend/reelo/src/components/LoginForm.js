import {Grid, TextField, Button} from '@material-ui/core';
import React from 'react';

export default function LoginForm(props) {
	const updateEmail = event => {
		props.onEmailChange(event.target.value);
	};

	const updatePassword = event => {
		props.onPasswordChange(event.target.value);
	};

	return (
		<form onSubmit={props.onSubmit}>
			<Grid container direction="column" alignItems="center" spacing={24}>
				<Grid item xs={12}>
					<TextField
						required
						type="email"
						label="Posta elettronica"
						value={props.emailValue}
						onChange={updateEmail}
					/>
				</Grid>
				<Grid item xs={12}>
					<TextField
						required
						type="password"
						label="Parola chiave"
						value={props.passwordValue}
						onChange={updatePassword}
					/>
				</Grid>

				<Grid item xs={12}>
					<Button type="submit" variant="contained" color="primary">
						Accedi
					</Button>
				</Grid>
			</Grid>
		</form>
	);
}
