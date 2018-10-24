## Challenge2 - Formatting Output and Data Types

### Objectives

The intent for this lesson is to build a more dynamic script that utilizes user input when ran. This added flexibility presents some challenges that we will need to overcome in the form of error handling and providing the user with helpful feedback.

Data Types
- Two very common data types are `dicts` and `lists`. You'll also see dicts with their `key:value` pairs where the `value` is actually a list itself. A `dict` of `lists`. Becoming familiar with this pattern will be very useful for manipulating industry standard outputs such as JSON.

Loops
- Mastering loops is integral to running scalable scripts. Becoming familiar with `for` and `while` loops will allow you to write scripts that iterate smartly.

Standardized Output
- Producing an output that adheres to an industry standard such as JSON for others to consume and manipulate is an important step in creating scripts.

---

### Instructions
Use [Star Wars API](https://www.swapi.co)
- API Request - Pull all starships from a specific film.
- Data Types - Create your own data objects to load API values in to.
- Loops - Load Starships in to your data objects.
- Loops - Load Pilots as a sub-object of their starships.
- Standardized Output - Output should be in JSON format.

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
