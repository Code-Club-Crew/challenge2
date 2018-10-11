## Challenge2 - Formatting Output

Using the [Star Wars API](https://www.swapi.co) pick a film and list only the starships with pilots flown in the film as well as the pilots of those star ships.

Output should be in JSON format

---
### Usage
- Clone this repo to your local computer using `git clone https://github.com/Code-Club-Crew/challenge2.git`.
- Create a new branch using `git checkout -b <NEW-BRANCH>`.
- Add your own subfolder to the repo.
- Add your files to the subfolder.

---

### Language
User choice

Recommend using the same language as Challenge1, Go, or Python

---
### Additional Resources
https://jsonlint.com/ - for checking the format of your output

https://swapi.co/documentation - for documentation relating to SWAPI

Check your language of choice for a supporting library.

---
### Example Output
```
{
  "ships": [
    {
      "name": "Millennium Falcon",
      "pilot": [
        {
          "name": "Chewbacca"
        },
        {
          "name": "Han Solo"
        },
        {
          "name": "Lando Calrissian"
        },
        {
          "name": "Nien Nunb"
        }
      ]
    },
    {
      "name": "X-wing",
      "pilot": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Biggs Darklighter"
        },
        {
          "name": "Wedge Antilles"
        },
        {
          "name": "Jek Tono Porkins"
        }
      ]
    },
    {
      "name": "Slave 1",
      "pilot": [
        {
          "name": "Boba Fett"
        }
      ]
    },
    {
      "name": "Imperial shuttle",
      "pilot": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Chewbacca"
        },
        {
          "name": "Han Solo"
        }
      ]
    }
  ],
  "name": "The Empire Strikes Back"
}
```
