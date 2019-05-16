import {createMuiTheme} from '@material-ui/core/styles';

import red from '@material-ui/core/colors/red';

const theme = createMuiTheme({
	typography: {
		useNextVariants: true,
	},
	palette: {
		primary: {
			// main: '#f5f5f5',
			main: '#b39ddb',
		},
		secondary: {
			// main: '#f5f5f5',
			main: '#f5f5f5',
		},

		error: red,
		contrastThreshold: 3,
		tonalOffset: 0.2,
	},

	props: {
		Link: {
			underlined: 'never',
		},

		MuiButton: {},

		MuiCard: {
			elevation: 0,
		},
	},

	overrides: {
		Link: {
			textDecoration: 'none',
		},
	},
});
theme.overrides.MuiCard = {};
export default theme;
