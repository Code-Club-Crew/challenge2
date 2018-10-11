#! /usr/bin/python

import swapi
import json

shiplist = []
pilotlist = []

filmdict = {
    "name": '',
    "ships": shiplist
    }

film = swapi.get_film(2)
shipindex = 0
filmtitle = film.title.encode('utf-8')
filmdict['name'] = filmtitle
film_ships = film.get_starships()
for ship in film_ships.iter():
    shipdict = {
        "name": '',
        "pilot": pilotlist
        }
    pilotlist = []
    pilotindex = 0
    shipdict['name'] = ship.name.encode('utf-8')
    ship_pilots = ship.get_pilots()
    if ship_pilots.count() > 0:
        filmdict['ships'].append(shipdict.copy())
        for pilot in ship_pilots.iter():
            pilotdict = {
                "name": ''
                }
            pilotdict['name'] = pilot.name.encode('utf-8')
            filmdict['ships'][shipindex]['pilot'].append(pilotdict.copy())
            pilotindex = pilotindex + 1
        shipindex = shipindex + 1
print(json.dumps(filmdict))
