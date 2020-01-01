import React, {useState} from 'react';
import {Grid, Typography, TextField, Paper, Button} from '@material-ui/core';

export const NamesakeForm = props => {
	const namesake = props.namesake;

	const [solver, setSolver] = useState([...namesake.solver]);
	const [comment, setComment] = useState(namesake.comment);
	const handleChange = e => setComment(e.target.value);

	const handleAccept = () => props.onAccept(namesake);
	const handleComment = () => props.onComment(namesake, comment);
	const handleMerge = () =>
		props.onMerge(next => {
			if (next && next.playerID === namesake.playerID) {
				const tmp = [...solver, ...next.solver];
				setSolver(tmp);
				namesake.solver = tmp;
				return namesake;
			} else {
				// TODO: this is lazy, i have no time to do fancier stuff
				window.alert('Unione non valida, nulla è cambiato');
			}
		});

	const renderSolver = solver =>
		solver.map((s, i) => (
			<Grid item xs={4} key={i}>
				<Paper>
					<Typography align="center">
						{s.City || 'Nessuna sede trovata'}
					</Typography>
					<Typography align="center">{s.Category}</Typography>
					<Typography align="center">
						{s.IsParis ? 'Intern' : 'N'}azionale
					</Typography>
					<Typography align="center">{s.Year}</Typography>
				</Paper>
			</Grid>
		));

	const renderActions = () => (
		<Grid item container justify="space-evenly" spacing={8} xs={3}>
			<Grid item xs={12} lg={5}>
				<Button color="primary" onClick={handleAccept} variant="contained">
					Accetta
				</Button>
			</Grid>
			<Grid item xs={12} lg={5}>
				<Button color="primary" onClick={handleComment} variant="contained">
					Commenta
				</Button>
			</Grid>
			<Grid item xs={12} lg={2}>
				<Button color="secondary" onClick={handleMerge} variant="contained">
					Unisci
				</Button>
			</Grid>
		</Grid>
	);

	return (
		<Grid item container spacing={8} xs={12}>
			<Grid item xs={1}>
				<Typography align="left">
					{`${namesake.name} ${namesake.surname} (${namesake.id})`}
				</Typography>
			</Grid>
			<Grid item container spacing={16} xs={6}>
				{renderSolver(namesake.solver)}
			</Grid>
			<Grid item xs={2}>
				<TextField
					variant="outlined"
					value={comment}
					onChange={handleChange}
					multiline
					rows={3}
				/>
			</Grid>
			{renderActions()}
			<Grid item xs={12}>
				<hr />
			</Grid>
		</Grid>
	);
};
