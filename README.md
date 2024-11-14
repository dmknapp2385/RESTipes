# Recipe Book

![Go](https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge)

# Description

Recipe book is a simple MVC using go(lang). The Gin Webframe work was used as a server to create RESTful API endpoints. Go's httpClient was used to create requests and a simple command line interface was used as the view. This code could easily be used as the back end to a web application by replacing the current view file with a front end.

## Table of Contents

- [Technologies](#technologies)
- [Usage](#usage)
- [Credits](#credits)

## Technologies used

- go
- Gin
- httpClient

## Usage

Follow the steps below to navigate and use the Recipe Book application:

1. Ensure you have go installed on your machine.
2. download all files into your favorite editor (VS Code was used in this production)
3. In the main package type go build into the command line.
4. Execute the executable file by running ./recipies.exe in the command line. Follow the prompts on the screen.
5. The application has the option of uploading a json file of Recipies on startup of the file using a command line flag. To do this: run the executable with the flag -file Example: ./recipies.exe -file=example.json. Ensure that the json file is within the same main directory as the executable file.
6. For creating the json file, follow the following recipe struct
   {
   "title":"your title",
   "author":"author",
   "ingredients":["a", "list", "of", "ingredients"],
   "steps" : ["a", "list", "of", "steps"],
   "baketime" : 120,
   "rating": 5
   }
7. Currently the application is set up so that if a recipe is created with a title that already exists, the new recipe will override the current recipe.

## Credits

[Elle Knapp](https://github.com/dmknapp2385)  
[Jeffry Freeman]()  
[Colton ]()
