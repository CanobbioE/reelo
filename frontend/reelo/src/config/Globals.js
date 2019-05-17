export default class Globals {
	static port = '8080';
	static baseURL = `${window.location.protocol}//${window.location.hostname}:${
		Globals.port
	}`;

	// All the APIs endpoints
	static API = {
		login: '/admin',
		upload: '/upload',
		ranks: '/ranks',
		force: '/force-reelo',
	};

	// All the routes in this app
	static routes = {
		home: '/',
		about: '/informazioni',
		upload: '/carica',
		admin: '/amministrazione',
		varchange: '/modifica-algoritmo',
	};
}
