# api client

using leaguepedia api

saving data in json which is used in api

### bot password
to setup for yourself, you have to make a bot password

use this documentation: https://lol.fandom.com/wiki/Help:Leaguepedia_API

### data path
when using outside of docker container, setup a .env file

this file and all the variables inside will be read in main.py with python-dotenv

add a variable inside the .env file called DATA_PATH, and give it a value of the path you want to save json data
