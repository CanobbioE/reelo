import React from 'react';
import {Grid, CircularProgress} from '@material-ui/core/';

export default props => {
	let ret = props.show ? (
		<Grid item>
			<br />
			<CircularProgress disableShrink />
			<br />
		</Grid>
	) : (
		''
	);
	return ret;
};
