import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import {Link} from 'react-router-dom';
import Globals from '../../config/Globals';

function DrawerList(props) {
	return (
		<List>
			<ListItem button component={Link} to={Globals.routes.home} key={'home'}>
				<ListItemText primary={'Inizio'} />
			</ListItem>

			<ListItem button component={Link} to={Globals.routes.ranks} key={'ranks'}>
				<ListItemText primary={'Classifiche'} />
			</ListItem>

			<ListItem
				button
				component={Link}
				to={Globals.routes.upload}
				key={'upload'}>
				<ListItemText primary={'Upload'} />
			</ListItem>
		</List>
	);
}
export default DrawerList;
