import React from 'react';
import {Button, Dialog, DialogActions, DialogTitle} from '@material-ui/core';

export default function DialogAlert(props) {
	return (
		<Dialog
			open={props.open}
			onClose={props.onClose}
			aria-labelledby="alert-dialog-title"
			aria-describedby="alert-dialog-description">
			<DialogTitle id="alert-dialog-title">
				Sono stati riscontrati errori nel documento
			</DialogTitle>

			{props.content}
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
}
