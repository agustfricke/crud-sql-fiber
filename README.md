# GO + MYSQL CRUD 

## Usage


- Clone the repo
```bash
git clone https://github.com/agustfricke/crud-sql-fiber
```

- Run the database
```bash
sudo docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:8.0-debian
sudo docker exec -it some-mysql bash
mysql -u root -p
create database albums;
use albums;
```

- Create the tables
```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);
```

- Run the server
```bash
cd crud-sql-fiber
go run main.go
```

### Test the endpoints

- Create
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "title": "Hola Mundo!",
  "artist": "Agustin",
  "price": 69.99
}' http://localhost:6969/add
```

- Get all
```bash
curl http://localhost:6969/all
```

- Get by ID
```bash
curl http://localhost:6969/id/1
```

- Get by artist name
```bash
curl http://localhost:6969/name/Agustin
```

- Delete by ID
```bash
curl -X DELETE http://localhost:6969/delete/1
```

- Update by ID
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
  "title": "edit",
  "artist": "edit",
  "price": 420
}' http://localhost:6969/put/1
```
