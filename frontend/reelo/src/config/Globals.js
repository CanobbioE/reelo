export default class Globals {
	static port = '8080';
	static baseURL = `${window.location.protocol}//${window.location.hostname}:${Globals.port}`;

	// All the APIs endpoints
	static API = {
		login: '/admin',
		upload: '/upload',
		ranks: '/ranks',
		force: '/force-reelo',
		algorithm: '/algorithm',
		count: '/count',
		years: '/years',
		exist: '/upload/exist',
		namesakes: '/namesakes',
	};

	// All the routes in this app
	static routes = {
		home: '/',
		about: '/informazioni',
		upload: '/carica',
		admin: '/amministrazione',
		varchange: '/modifica-algoritmo',
		namesakes: '/omonimi',
	};
}
