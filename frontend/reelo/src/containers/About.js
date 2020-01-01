import {Grid, Typography} from '@material-ui/core';
import React from 'react';

import {withStyles} from '@material-ui/core/styles';

const styles = () => ({
	title: {
		marginTop: '28px',
		marginBottom: '15px',
	},
});

function About(props) {
	const {classes} = props;
	return (
		<Grid container justify="center">
			<Grid item xs={10}>
				<Typography variant="h4" className={classes.title}>
					Reelo
				</Typography>
				<Typography variant="subtitle1">
					<em>
						La classifica REELO nasce con l'obiettivo di creare una graduatoria
						per i giochi matematici che tenga conto dei risultati di ogni
						concorrente attraverso gli anni, in analogia alla classifica
						scacchistica ELO e alla classifica tennistica ATP. Il nome REELO
						vuole richiamare il sistema Elo e in Esperanto significa "numero
						reale".
					</em>
				</Typography>

				<dl>
					<dt>
						<strong> Ideazione:</strong>
					</dt>
					<dd> Cesco Reale.</dd>
					<dt>
						<strong>Comitato scientifico: </strong>
					</dt>
					<dd>Marco Broglia, Andrea Nari, Marco Pellegrini, Cesco Reale.</dd>
					<dt>
						<strong>Implementazione: </strong>
					</dt>
					<dd>Edoardo Canobbio, in collaborazione con Anna Bernardi. </dd>
					<dt>
						<strong>Comitato tecnico: </strong>
					</dt>
					<dd>
						Fabio Angelini, Anna Bernardi, Edoardo Canobbio, Mirko Cappuccia.
					</dd>
					<dt>
						<strong>Consulenti comitato scientifico: </strong>
					</dt>
					<dd>
						David Barbato, Maurizio De Leo, Francesco Morandin, Alberto Saracco.
					</dd>
					<dt>
						<strong>Consulenti comitato tecnico: </strong>
					</dt>
					<dd>
						Diego Corsatto, Carlo Frittoli, Alessio Palmero, Elia Scandaletti
					</dd>
				</dl>
			</Grid>
		</Grid>
	);
}

export default withStyles(styles)(About);
