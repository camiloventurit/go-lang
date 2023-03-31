
# Movies

This task requires the development of an API with a specific functionality related to movie search. The API will enable users to search for movies by title using an exact value that is passed in as an argument.

The API will have two different modes of operation depending on the availability of the movie in the local database. If the exact match for the movie title exists in the local database, the API will retrieve the details of that movie and return them to the user. However, if there is no exact match in the local database, the API will use the imdb-api package to search for the movie on the Internet Movie Database (IMDb).

If the IMDb search returns one or more results, the API will store the details of the first result in the local database and return the details to the user. However, if the IMDb search does not return any results, the API will return a message indicating that the movie was not found.

Overall, this API will enable users to search for movies using an exact title match and will use IMDb as a fallback option if the movie is not found in the local database.

**by ID**

```
http://localhost:8080/api/search-by-title?queryType=i&title=tt0120338
```

**by Title**\
where queryType could be "t" or "i"

```
http://localhost:8080/api/search-by-title?queryType=t&title=Birds of prey
```

t = title\
i = id
