# goAvitoDB

### Launch

Create an .env file like this in the root directory of your project:

``` .env
DB_PASSWORD=...

PGADMIN_DEFAULT_EMAIL=...
PGADMIN_DEFAULT_PASSWORD=...
```
<details>
  <summary>Entity-relationship diagram</summary>
  <p align="center">
    <img src=diagram/diagram.png width=50% />
  </p>
</details>

Enter this command in a terminal running Docker:

```
docker-compose up --build
```

In order to check the database operation go to http://localhost:5050
