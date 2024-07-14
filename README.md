# URL-Shortener

Made simply and with few functions, the project consists of shortening a long URL submitted by the user. Furthermore, this repository contains solution design, ORM mapping, use of third-party libraries, use of GIT and creativity.

The project consist in two cases:
1. Shorten URL
2. Retrieve URL

## Shorten URL
![image](https://github.com/user-attachments/assets/82285503-f513-4531-9a1e-6591f815ba41)

1. User submits a long URL
2. API generates a short URL by randomizing characters in the range [A-Z], [a-z], [0-9], bringing the possibility of 62^7 shortened URLs
3. Checks whether the long URL exists in cache (we can use Redis, for example). If yes, retrieve it from the cache. If not, add the two URLs (long and short) and their respective hash to the database.
4. Returns the shortened URL to the user

## Retrieve URL

## Things to consider
