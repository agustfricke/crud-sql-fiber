# GO + MYSQL CRUD 

```bash
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8.0-debian
sudo docker exec -it some-mysql bash
mysql -u root -p
create database todo;
use todo;
```

```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

## Usage
- Run the scripts inside of the folder **requests**

```bash
# post request
./requests/post
```
