# reelo

Reelo is an ELO system for Mathematics games. The name refers to the Esperanto's term that means "real number".

## TODO

### Back end

- middleware for authentication
- prompt for error while parsing
- parsing files form user input and not from whole folder
- define APIs entities
- allow ELO algorithm to be edited -> needs a table in the db
- endpoint for ranks
- Actual ELO algorithm

### Front end

- Edit algorithm form
- Upload files form
- Display errors (especially while parsing files)
- Better cookies handling
- Rankings page

## Application flow

There's two types of user: Admin and Standard.
**Admin** needs to login -> Admin loads a ranking file -> file get parsed -> if parsing error Admin will be prompted to fix it -> parsed file gets converted into entity -> entity goes into db -> after all entities for the file are in the db Reelo gets (re)calculated for al the entities.

**Admin** needs to login -> Admin changes variables -> db gets updated -> Reelo gets recalculated for all entities.

**User** will visit the rankings page -> server returns a list of all the entities and their Reelo -> User is happy =).

## Credits

- Ideation: Cesco Reale
- Implementation: Fabio Angelini
- Scientific Committee: Fabio Angelini, Andrea Nari, Marco Pellegrini, Cesco Reale
- Collaborators: David Barbato, Maurizio De Leo, Francesco Morandin, Alberto Saracco
