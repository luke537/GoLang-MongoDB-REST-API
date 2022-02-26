# GoLang-MongoDB-REST-API by Luke Ryan

This is a REST API written in GoLang, that connects to a MongoDB database cluster, providing full CRUD functionality.

The MongoDB database contains data about bins (waste containers), with the following attributes:
1. ID
2. Name
3. Longitude/Latitude
4. Accepted Materials


The API has the following endpoint functions:
1. Get all bins within a specified distance (in metres) and that contain a given waste material
2. Add a new bin to the database with a given name, longitude/latitude and waste material
3. Update an existing bin, given its ID
4. Delete an existing bin, given its ID

# Usage Examples
