import React, {useEffect} from 'react';
import Globals from '../config/Globals';
import connect from 'react-redux';

export default ChildComponent => {
	const ComposedComponent = props => {
		useEffect(() => {
			shouldNavigateAway();
		}, []);

		const shouldNavigateAway = () => {
			if (!this.props.auth) {
				this.props.history.push(Globals.routes.home);
			}
		};

		return <ChildComponent {...props} />;
	};
	function mapStateToProps(state) {
		return {
			auth: state.auth.authenticated,
		};
	}
	return connect(mapStateToProps)(ComposedComponent);
};
