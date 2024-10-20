# pokedex-cli

You need Go 1.23.1

## Ways to run the program
### Method 1
Download the latest release from the releases section: [Releases](https://github.com/navivan123/pokedex-cli/releases)
Run the desired binary depending on your operating system.

### Method 2
If you have go 1.23.1, you can simply run the CLI on Linux using the run.sh script.
You can also use whatever go commands to build and run the program.

## CLI Breakdown:

### help
Lists the various commands you can execute on the cli

### exit
Exit the program.  Alternatively use Ctrl+C or whatever interrupt signal can kill it.

### map
When first ran, lists the first 20 locations of the pokemon world. Use map again to view the *next* 20 locations. If the last page is already selected, it will display a message that it is at the end of the list and re-display the last few remaining locations.  The last few remaining locations are not necessarily 20.

### mapb
When first ran, lists the first 20 locations of the pokemon world. Use mapb again to view the *previous* 20 locations. If the first page is already selected, it will display a message that it is at the start of the list and re-display the first 20 locations.

### explore
Use a location available from the list of locations displayed on map/mapb to view all the pokemon available in the area. You need to feed it a second argument, the area (it will ignore any more arguments) Ex: `explore canalave-city-area`.

### catch
Use a pokemon's name to try and catch the pokemon. Ex: `catch pikachu`. You have a certain probability of catching said pokemon, based on the pokemon's base XP. This isn't how pokemon works, but it makes my job way simpler, so we'll do it like that.

### inspect
View a certain pokemon's stats based on the name you select. Ex: `inspect pikachu`. You can only view the pokemon's stat if you have caught this pokemon and is residing on your pokedex, otherwise the program will tell you and give you no information about that pokemon. 

### pokedex
View all the pokemon you've caught.  Pretty self-explanatory, only lists the names of the pokemon.  If you want to look at the stats of a pokemon you have caught, use inspect.

## Advanced Info:
- If you know the ID of the pokemon/area, you can also use that.
- Each pokemons Base XP's square root is taken and then compared with a uniform distribution between 0 and sqrt(800). This produces a curve that works nicely, where 50%/50% chance coincides with an XP of 200, so unless the pokemon has a high base XP, it should be easily captured, though the probability is not 100.

