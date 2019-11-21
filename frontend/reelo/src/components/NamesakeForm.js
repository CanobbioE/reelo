import React, {useState} from 'react';
import {Grid, Typography, TextField, Paper, Button} from '@material-ui/core';

export const NamesakeForm = props => {
	const namesake = props.namesake;

	const [solver, setSolver] = useState([...namesake.solver]);
	const [comment, setComment] = useState('');
	const handleChange = e => setComment(e.target.value);

	const handleAccept = () => props.onAccept(solver);
	const handleComment = () => props.onComment(comment);
	const handleMerge = () =>
		props.onMerge(async next => {
			if (next && next.playerID === namesake.playerID) {
				const tmp = [...solver, ...next.solver];
				await setSolver(tmp);
				console.log('set to', tmp);
				return [tmp, null];
			}
			return [null, 'Error'];
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
				<Button color="secondary" onClick={handleComment} variant="contained">
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
