## Challenge1
Using the [Star Wars API](https://www.swapi.co) for each film list the starships used in the films as well as the pilots of those star ships.
Output should be in JSON format
---
### Usage
- Clone this repo to your local computer using `git clone https://github.com/Code-Club-Crew/challenge1.git`.
- Create a new branch using `git checkout -b <NEW-BRANCH>`.
- Add your own subfolder to the repo.
- Add your files to the subfolder.

---

### Language
User choice

Recommend Go or Python

---
### Additional Resources
https://jsonlint.com/ - for checking the format of your output
https://swapi.co/documentation

Check your language of choice for a supporting library.

---
### Example Output
```
{
  "film": [
    {
      "name": "Film1",
      "shipsandpilots": [
        {
        "starship": "starship1",
        "pilots": [
          "pilot1",
          "pilot2"    
          ]
        },  {
        "starship": "starship2",
        "pilots": [
          "pilot3",
          "pilot4"    
          ]
        }
      ]
    },  {
      "name": "Film2",
      "shipsandpilots": [
        {
        "starship": "starship3",
        "pilots": [
          "pilot5",
          "pilot6"    
          ]
        }, {
        "starship": "starship4",
        "pilots": [
          "pilot7",
          "pilot8"    
          ]
        }
      ]
    }
  ]
}
```
