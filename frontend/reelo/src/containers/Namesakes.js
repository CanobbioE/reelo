import React, {useEffect, useState} from 'react';
import {connect} from 'react-redux';
import {compose} from 'redux';
import {Grid, Typography, Button} from '@material-ui/core';
import {withStyles} from '@material-ui/core/styles';
import {
	fetchNamesakes,
	updateNamesake,
	acceptNamesake,
	commentNamesake,
} from '../actions';
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
	const [page, setPage] = useState(1);
	useEffect(() => {
		props.fetchNamesakes(page, 300);
	}, [page]);

	const handleMerge = i => merge => {
		const ret = merge(props.analysis.namesakes[i + 1]);
		if (ret) {
			props.updateNamesake(i, ret);
		}
	};

	const handleAccept = async namesake => {
		if (!window.confirm('Sicuro di voler confermare la soluzione proposta?')) {
			return;
		}
		await props.acceptNamesake(namesake);
		props.fetchNamesakes(page, 300);
		if (props.analysis.error && props.analysis.error !== '') {
			window.alert(props.analysis.error);
		}
	};

	const handleComment = async (namesake, comment) => {
		await props.commentNamesake(namesake, comment);
		props.fetchNamesakes(page, 300);
	};

	const renderNamesakes = () =>
		props.analysis &&
		props.analysis.namesakes &&
		props.analysis.namesakes.map((namesake, i) => (
			<NamesakeForm
				key={`${namesake.playerID} ${namesake.id}`}
				namesake={namesake}
				onComment={handleComment}
				onAccept={handleAccept}
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
			<Grid item xs={11}>
				<Typography variant="subtitle1">
					Utilizza questa pagina per risolvere i casi di omonimia. Ci sono tre
					azioni disponibili:
					<dl>
						<dt>
							<b>Accetta:</b>
						</dt>
						<dd>
							Conferma la soluzione proposta e crea un nuovo giocatore con la
							cronologia di partecipazioni proposta.
						</dd>

						<dt>
							<b>Unisci:</b>
						</dt>
						<dd>
							La cronologia del giocatore selezionato viene unita alla
							cronologia del suo omonimo sottostante. Se i due giocatori non
							sono omonimi non succede nulla. La modifica viene attuata solo una
							volta premuto il pulsante "Accetta".
						</dd>

						<dt>
							<b>Commenta:</b>
						</dt>
						<dd>
							Semplicemente aggiunge un commento in modo da poi risolvere
							manualmente l'omonimia. Nessuna modifica drastica viene applicata.
							Non necessita che venga premuto il tasto "Accetta".
						</dd>
					</dl>
					Per questioni di performanza i giocatori vengono analizzati a pagine
					di 300 giocatori alla volta. E l'ordine in cui compaiono non &egrave;
					costante.
					<br />
					<br />
					<Grid item container spacing={8} justify="center">
						<Grid item xs={12}>
							<Typography align="center">
								Al momento sei a pagina {page}
							</Typography>
						</Grid>
						<Grid item>
							<Button variant="contained" onClick={() => setPage(page - 1)}>
								Diminuisci pagina
							</Button>
						</Grid>
						<Grid item>
							<Button variant="contained" onClick={() => setPage(page + 1)}>
								Aumenta pagina
							</Button>
						</Grid>
					</Grid>
				</Typography>
				<hr />
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
	connect(mapStateToProps, {
		fetchNamesakes,
		updateNamesake,
		acceptNamesake,
		commentNamesake,
	}),
);

export default withStyles(styles)(composedComponent(Namesakes));
