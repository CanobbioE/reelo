export default class Globals {
	static baseURL = 'http://localhost:8080';

	// All the APIs endpoints
	static API = {
		login: '/login',
	};

	// All the routes in this app
	static routes = {
		home: '/',
		ranks: '/classifiche',
		upload: '/carica',
		admin: '/amministrazione',
	};
}
