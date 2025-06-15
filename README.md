# Bristol Bites 🍽️

A lightweight Go application for storing and retrieving information about restaurants, cafés, and food spots around Bristol.

## Purpose

This personal project aims to:

- Maintain a simple, local database of eateries in Bristol.
- Store useful metadata like ratings, dietary options (vegan, gluten-free, etc), cuisine types, and more.
- Provide recommendations based on user preferences like dietary needs, rating thresholds, or mood.

## Features (Planned or In Progress)

- Add and edit restaurant/café entries (admins only?)
- Filter/search by dietary requirements, rating and proximity to a given location
- Where feasible, use API to check booking availablility
- Input validation
- Authentication middleware

## Why?

I wanted a simple, customizable way to track and explore Bristol’s food scene, tailored to my taste and needs, without relying on big or paid platforms.

## How to run locally

```
docker-compose up --build
```

OR 

```
make up
````