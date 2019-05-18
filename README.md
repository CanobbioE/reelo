# Reelo

Reelo is an ELO system for Mathematics games. The name refers to the Esperanto's term that means "real number".

## TODO

### Back end

- Handle double surnames !!
- API limiter (?)
- Refactor db
- Better error handling
- Better auth handling
- Pagination (?)

### Front end

- i18n (?)
- Caching
- Better cookies handling
- A bit of refactoring wouldn't hurt

## Application flow

There's two types of user: Admin and Standard.

**Admin** needs to login -> Admin loads a ranking file -> file get parsed -> if parsing error Admin will be prompted to fix it -> parsed file gets converted into entity -> entity goes into db -> after all entities for the file are in the db Reelo gets (re)calculated for al the entities.

**Admin** needs to login -> Admin changes variables -> db gets updated -> Reelo gets recalculated for all entities.

**User** will visit the rankings page -> server returns a list of all the entities and their Reelo -> User is happy =).

## Credits

- Ideation: Cesco Reale
- Implementation: Fabio Angelini, Anna Bernardi, Edoardo Canobbio
- Scientific Committee: Fabio Angelini, Andrea Nari, Marco Pellegrini, Cesco Reale
- Collaborators: David Barbato, Maurizio De Leo, Francesco Morandin, Alberto Saracco
