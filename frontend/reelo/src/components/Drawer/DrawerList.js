import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import {Link} from 'react-router-dom';
import Globals from '../../config/Globals';

function DrawerList(props) {
	return (
		<List>
			<ListItem button component={Link} to={Globals.routes.home} key={'ranks'}>
				<ListItemText primary={'Classifiche'} />
			</ListItem>

			{props.isAuthenticated && (
				<ListItem
					button
					component={Link}
					to={Globals.routes.upload}
					key={'upload'}>
					<ListItemText primary={'Caricamento'} />
				</ListItem>
			)}

			{props.isAuthenticated && (
				<ListItem
					button
					component={Link}
					to={Globals.routes.varchange}
					key={'varchange'}>
					<ListItemText primary={'Modifica algoritmo'} />
				</ListItem>
			)}

			<ListItem button component={Link} to={Globals.routes.about} key={'home'}>
				<ListItemText primary={'Informazioni'} />
			</ListItem>
		</List>
	);
}
export default DrawerList;
