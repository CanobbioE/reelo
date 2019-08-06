import {Button} from '@material-ui/core';
import React from 'react';

export default function Logout(props) {
	return (
		<Button variant="outlined" onClick={props.onClick}>
			Esci
		</Button>
	);
}
