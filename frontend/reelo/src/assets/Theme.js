// Author: Edoardo Canobbio
import {createMuiTheme} from '@material-ui/core/styles';

import red from '@material-ui/core/colors/red';

const theme = createMuiTheme({
	typography: {
		useNextVariants: true,
	},
	palette: {
		primary: {
			main: '#3A405A',
		},
		secondary: {
			main: '#FAFAFA',
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
