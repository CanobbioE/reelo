package parse

var commonSurnamePrefix = []string{
	"da",
	"dal",
	"dalla",
	"dalle",
	"de",
	"del",
	"della",
	"delle",
	"delli",
	"dello",
	//"dei",
	"di",
	"la",
	"le",
	"lo",
}

var doubleWordNames = []string{
	"alberto maria",
	"alioshka michele",
	"aly sonia",
	"andrea alberto",
	"andrea c.",
	"andrea carlo",
	"andrea celeste",
	"andrea clarissa",
	"andrea f.",
	"angela t.",
	"anna giulia",
	"anna laura",
	"anna m.",
	"anna maria",
	"anna paola",
	"anna sara",
	"anna vera",
	"anton giulio",
	"antonio francesco m.",
	"augusto francesco",
	"aurele renè",
	"carlo alberto",
	"carlo maria",
	"carlo s.",
	"cataldo luigi",
	"chiara m.",
	"claudio g.",
	"dario francesco",
	"davide c.",
	"diana andrea",
	"elena maria",
	"elena sofia",
	"elena tea",
	"elio virgilio",
	"erica m.",
	"ettore mario",
	"eugenio ernesto",
	"fabio massimo",
	"federica edvige",
	"federico ernesto",
	"federico flavio",
	"federico lorenzo",
	"federico mattia",
	"francesco g.",
	"francesco guido",
	"francesco saverio",
	"giacomo a.",
	"gian luca",
	"gian marco",
	"giovanna c.",
	"giovanni andrea",
	"giovanni b.",
	"giovanni battista",
	"giovanni luigi",
	"giulia a.",
	"giulia carlotta",
	"giulia m.",
	"giuliano andrea",
	"giuseppe p.",
	"giuseppe pio",
	"guido maria",
	"iris tullia",
	"jean pierre",
	"leone cesare",
	"linda anna",
	"lorenzo maria",
	"luca arduino",
	"luca donato",
	"luca felipe",
	"luca giulio",
	"lucio alberto",
	"luigi federico",
	"luigi tommaso",
	"m. chiara",
	"m. elena",
	"m. giovanna",
	"marco g.",
	"marco m. j.",
	"maria anna",
	"maria antonia",
	"maria assunta",
	"maria chiara",
	"maria cristina",
	"maria e.",
	"maria francesca",
	"maria giovanna",
	"maria giulia",
	"maria laura",
	"maria livia",
	"maria luisa",
	"maria rosa",
	"maria rosaria",
	"maria stella",
	"maria teresa",
	"maria vittoria",
	"marianne elena",
	"marie nicole",
	"marilena moira",
	"mario luca",
	"mario manuel",
	"mario o.",
	"marti'n",
	"matteo pio",
	"michael norbert",
	"micheal norbert",
	"michele luigi",
	"michele r.",
	"mihai codrut",
	"nicolao m.",
	"nicolo' f.",
	"paolo armando",
	"paolo emilio",
	"pier maria",
	"pio gabriele",
	"rita elisabeth",
	"roberta maria c.",
	"rosa m.gloria",
	"rosa maria",
	"sabrina sharon",
	"silvia anna",
	"sonia laura",
	"vito antonio",
	"vito gaetano",
	"alessandro mario",
	"alessandro pietro",
	"alessia f.",
	"alexandru maria",
	"ambra s.",
	"anna chiara",
	"anna flavia",
	"anna stella",
	"arturo federico",
	"berardo m.",
	"bl. marco",
	"c. filippo",
	"camillo aurelio",
	"carla m.",
	"carlo bruno",
	"carmelo ismaele",
	"cassandra bruna",
	"claudio filippo",
	"cosimo d.g.",
	"david antonio",
	"davide andrea",
	"davide pio",
	"elena luiz",
	"enrico gaetano",
	"enrico giuseppe",
	"enrico maria",
	"fausta pia",
	"filippo j.l.",
	"francesco fabio",
	"francesco maria",
	"gabriele marco",
	"gabriele pio",
	"george razvan",
	"gian franco",
	"gian michele",
	"gian nicola",
	"gianforte filippo",
	"giovanni agostino",
	"giovanni b.",
	"giovanni battista",
	"giulia maria",
	"jean manuel",
	"jean marie",
	"jounis jarir",
	"kamil karol",
	"kevin ludovico",
	"livio c.e.",
	"lorenzo ettore",
	"lorenzo piero",
	"loris maria",
	"luca pasquale",
	"m. antonietta",
	"m. claudi",
	"m. ilenia",
	"manfred marvin",
	"maria a.",
	"maria assunta",
	"maria elena",
	"maria g.",
	"maria giacinta",
	"maria grazia",
	"maria luigia",
	"maria nunzia",
	"maria pia ida",
	"maria pia",
	"maria rita",
	"maria sole",
	"marica o.",
	"martina fang",
	"matias fernando",
	"matteo giovanni",
	"matteo j.",
	"michela pia",
	"mohamed ismail",
	"nicola maria",
	"patricia diana",
	"pie r",
	"pier giorgio",
	"racek lucas",
	"s. emanuele",
	"salvatore l.",
	"silvia v.",
	"tatyana g.",
	"ten lon",
	"valentina rebecca",
	"valerio luca",
	"wen jinc",
	"wen qian",
	"younis jarir",
}
