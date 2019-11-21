import React, {useEffect} from 'react';
import {connect} from 'react-redux';
import {compose} from 'redux';
import {Grid, Typography} from '@material-ui/core';
import {withStyles} from '@material-ui/core/styles';
import {fetchNamesakes} from '../actions';
import {NamesakeForm} from '../components/NamesakeForm';
import LoadingIcon from '../components/LoadingIcon';

const styles = theme => ({
	title: {
		marginTop: '28px',
		marginBottom: '15px',
		paddingLeft: 5 * theme.spacing.unit,
	},
	subtitle: {
		marginTop: '28px',
		marginBottom: '15px',
	},
});

const Namesakes = props => {
	const {classes} = props;
	useEffect(() => {
		props.fetchNamesakes(1, 100);
	}, []);

	const handleMerge = i => async merge => {
		const [ret, err] = await merge(props.analysis.namesakes[i + 1]);
		if (err) {
			return;
		}
		props.analysis.namesakes[i] = ret;
		props.analysis.namesakes[i + 1] = null;
	};

	const renderNamesakes = () =>
		props.analysis &&
		props.analysis.namesakes &&
		props.analysis.namesakes.map((namesake, i) => (
			<NamesakeForm
				key={`${namesake.playerID} ${namesake.id}`}
				namesake={namesake}
				onComment={c => console.log('commented', c)}
				onAccept={solver => console.log('accepted', solver)}
				onMerge={handleMerge(i)}
			/>
		));

	const renderError = () => (
		<Grid item xs={12}>
			<Typography color="error" align="center">
				{props.analysis.error}
			</Typography>
		</Grid>
	);

	return (
		<Grid container justify="center">
			<Grid item xs={11}>
				<Typography variant="h4" className={classes.title}>
					Risoluzione Omonimi
				</Typography>
			</Grid>

			<Grid item container spacing={24} xs={10}>
				<Grid item xs={1}>
					<Typography className={classes.subtitle} align="left" variant="h6">
						Giocatore
					</Typography>
				</Grid>
				<Grid item xs={6}>
					<Typography className={classes.subtitle} align="center" variant="h6">
						Soluzione proposta
					</Typography>
				</Grid>
				<Grid item xs={2}>
					<Typography className={classes.subtitle} align="left" variant="h6">
						Commento
					</Typography>
				</Grid>
				<Grid item xs={3}>
					<Typography className={classes.subtitle} align="center" variant="h6">
						Azioni
					</Typography>
				</Grid>
				{props.analysis.loading && !props.analysis.error.length
					? null
					: renderNamesakes()}
				{props.analysis.error.length ? renderError() : null}
				<LoadingIcon show={props.analysis.loading} />
			</Grid>
		</Grid>
	);
};

function mapStateToProps({uploadForm, analysis}) {
	return {uploadForm, analysis};
}

const composedComponent = compose(
	// TODO RequireAuth,
	connect(
		mapStateToProps,
		{
			fetchNamesakes,
		},
	),
);

export default withStyles(styles)(composedComponent(Namesakes));
