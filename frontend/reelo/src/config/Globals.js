export default class Globals {
    static port = "8080";
    static baseURL = `${window.location.protocol}//${window.location.hostname}:${Globals.port}`;

    // All the APIs endpoints
    static API = {
        players: {
            count: "/players/count",
            reelo: {
                calculate: "/players/reelo/calculate",
            },
            comment: "/players/comment",
        },
        ranks: {
            all: "/ranks/all",
            upload: "/ranks/upload",
            exist: "/ranks/exist",
            years: "/ranks/years",
        },
        auth: {
            login: "/auth/login",
        },
        namesakes: {
            all: "/namesakes/all",
            update: "/namesakes/update",
        },
        costants: {
            all: "/costants/all",
            update: "/costants/update",
        },
    };

    // All the routes in this app
    static routes = {
        home: "/",
        about: "/informazioni",
        upload: "/carica",
        admin: "/amministrazione",
        varchange: "/modifica-algoritmo",
        namesakes: "/omonimi",
    };
}
