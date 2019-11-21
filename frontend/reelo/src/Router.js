import {MuiThemeProvider} from '@material-ui/core/styles';
import {BrowserRouter, Route} from 'react-router-dom';
import {createStore, applyMiddleware} from 'redux';
import React, {Component} from 'react';
import {Provider} from 'react-redux';
import {DndProvider} from 'react-dnd';
import HTML5Backend from 'react-dnd-html5-backend';
import reduxThunk from 'redux-thunk';

import NavBar from './components/NavBar';
import {properties} from './config/Properties';
import Globals from './config/Globals';
import reducers from './reducers';
import About from './containers/About';
import Ranks from './containers/Ranks';
import Admin from './containers/Admin';
import Upload from './containers/Upload';
import EditAlgorithm from './containers/EditAlgorithm';
import Namesakes from './containers/Namesakes';

import './index.css';

const store = createStore(
	reducers,
	{
		// A preloaded state
	},
	applyMiddleware(reduxThunk),
);

class Router extends Component {
	render() {
		return (
			<DndProvider backend={HTML5Backend}>
				<MuiThemeProvider theme={properties.theme}>
					<Provider className="bg-white" store={store}>
						<BrowserRouter>
							<NavBar />
							<div>
								<Route exact path={Globals.routes.home} component={Ranks} />
								<Route exact path={Globals.routes.about} component={About} />
								<Route exact path={Globals.routes.upload} component={Upload} />
								<Route exact path={Globals.routes.admin} component={Admin} />
								<Route exact path={Globals.routes.namesakes} component={Namesakes} />
								<Route exact path={Globals.routes.varchange} component={EditAlgorithm} />
							</div>
						</BrowserRouter>
					</Provider>
				</MuiThemeProvider>
			</DndProvider>
		);
	}
}

export default Router;
