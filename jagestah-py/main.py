#! /usr/bin/python

import swapi
import json

#Define empty variables as their appropriate type before use
shiplist = []
pilotlist = []
filmdict = {
    "name": '',
    "ships": shiplist
    }

#Get the information for The Empire Strikes Back
film = swapi.get_film(2)
#initialize shipindex for using as a reference for assign pilots to a specific ship
shipindex = 0
#Data from SWAPI sometimes needs to be encoded to deal with unicode
filmtitle = film.title.encode('utf-8')
filmdict['name'] = filmtitle
#get the list of ships from film(2)
film_ships = film.get_starships()
#Iterate over each spaceship returned by film.get_starships()
for ship in film_ships.iter():
    #Initialize the shipdict, pilotlist, and ppilotindex, so that they reset for each new ship
    shipdict = {
        "name": '',
        "pilot": pilotlist
        }
    pilotlist = []
    shipdict['name'] = ship.name.encode('utf-8')
    #get the pilots for each individual ship
    ship_pilots = ship.get_pilots()
    #Only add ships that have named pilots
    if ship_pilots.count() > 0:
        #Append ships to the shiplist list
        filmdict['ships'].append(shipdict.copy())
        #iterate over the pilots returned by ship.get_pilots()
        for pilot in ship_pilots.iter():
            #Initialize the variables so they're reset for each pilot
            pilotdict = {
                "name": ''
                }
            pilotdict['name'] = pilot.name.encode('utf-8')
            filmdict['ships'][shipindex]['pilot'].append(pilotdict.copy())
        #Increment the shipindex, used for assigning the pilots to the correct ship
        shipindex = shipindex + 1
#Convert all the python types to JSON equivalents and print the results
print(json.dumps(filmdict))
