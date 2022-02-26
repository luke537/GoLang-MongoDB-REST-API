# GoLang-MongoDB-REST-API by Luke Ryan

## Overview
This is a REST API written in GoLang, that connects to a MongoDB database cluster, providing full CRUD functionality.

The MongoDB database contains data about bins (waste containers), with the following attributes:

| Attribute | Type |
| ID                      | string |
| Name                    | string |
| Longitude               | float64 |
| Latitude                | float64 |
| Accepted Materials      | string[] |

## API Functions
- Get all bins within a specified distance (in metres) and that contain a given waste material
- Add a new bin to the database with a given name, longitude/latitude and waste material
- Update an existing bin, given its ID
- Delete an existing bin, given its ID

## Usage Examples

### Getting bins with given distance and given material
**HTTP Method:** GET  
**Parameters:** None
**URL:** /api/bins/
**Request Body:**
```json

{
    "distance": 3000, // In metres
    "longitude": -9.7263177,
    "latitude": 52.2835895,
    "material": "plastic"
}

```

### Adding a new bin with a given location and accepted materials
**HTTP Method:** POST
**Parameters:** None
**URL:** /api/bins/
**Request Body:**
```json

{
    "name": "Bin outside Centra",
    "longitude": -8.4779364,
    "latitude": 51.8981547,
    "acceptedMaterials": ["plastic", "cans"]
}

```

### Updating an existing bin
**HTTP Method:** PUT
**Parameters:** id
**URL:** /api/bins/{id}
**Request Body:**
```json

{
    "name": "Bin outside Centra",
    "longitude": -8.4779364,
    "latitude": 51.8981547,
    "acceptedMaterials": ["plastic", "cans"]
}

```

### Deleting an existing bin
**HTTP Method:** DELETE
**Parameters:** id
**URL:** /api/bins/{id}
**Request Body:**
```json

{
    "acceptedMaterials": ["glass"]
}

```
