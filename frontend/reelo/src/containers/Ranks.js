import {Grid, Typography, Button} from '@material-ui/core';
import React, {useEffect} from 'react';
import RanksTable from '../components/RanksTable';
import {connect} from 'react-redux';
import {compose} from 'redux';
import {fetchRanks, forceReelo} from '../actions';
import LoadingIcon from '../components/LoadingIcon';

function Ranks(props) {
	useEffect(() => {
		props.fetchRanks();
	}, []);

	const rows = props.ranks.rows;
	const labels = ['#', 'Nome', 'Cognome', 'Categoria', 'Reelo'];

	const content = (
		<Grid container item xs={12} justify="space-around">
			<Grid item xs={1}>
				<LoadingIcon show={props.ranks.loading} />
			</Grid>
			<Grid item xs={12}>
				{!props.ranks.loading && <RanksTable rows={rows} labels={labels} />}
			</Grid>
		</Grid>
	);

	const error = (
		<Grid item xs={12}>
			<Typography align="center" variant="body2" color="error">
				Oops, si &egrave; verificato un errore:{' '}
				{rows
					? props.ranks.error
					: 'Non sono presenti valori nella base di dati'}
			</Typography>
		</Grid>
	);
	return (
		<Grid container justify="center">
			<Grid item container spacing={24} xs={10}>
				<Grid item xs={12}>
					<Typography variant="h4">Classifiche</Typography>
				</Grid>
				{props.ranks.error === '' && rows ? content : error}
				<Grid item xs={12}>
					{!props.auth.authenticated ? null : (
						<Button
							onClick={() => {
								props.forceReelo();
								props.fetchRanks();
							}}
							variant="contained"
							color="primary">
							Ricalcola REELO
						</Button>
					)}
				</Grid>
			</Grid>
		</Grid>
	);
}

function mapStateToProps({ranks, auth}) {
	return {ranks, auth};
}

const composedComponent = compose(
	connect(
		mapStateToProps,
		{fetchRanks, forceReelo},
	),
);

export default composedComponent(Ranks);
