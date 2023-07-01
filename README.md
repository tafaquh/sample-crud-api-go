# Golang Simple CRUD
disclaimer: this is for learning purpose only. please do not make this repository structure as reference for your real project.

## How to run
```
$ go run main.go
```

## Setup
Please setup MYSQL database on your local machine and put the credentials into the .env file.
After you finish setup the MYSQL database, please open the mysql
```
~/go/src/github.com/tafaquh/coba-go                                                                                                                 11:29:54
> mysql -h 127.0.0.1 -uroot -p

```

and then do the followings command and costumize based your own needs
```
mysql> create database coba_go;
Query OK, 1 row affected (0.01 sec)

mysql> create user 'taf'@'%' identified by 'taf1234';
Query OK, 0 rows affected (0.01 sec)

mysql> grant all privileges ON *.* TO 'taf'@'%';
Query OK, 0 rows affected (0.01 sec)
```

## Main Features
1. Register User
2. Login User
3. CRUD Course

## What's next?
1. Please add JWT auth middleware to the router
2. Please try to limit some endpoints using the JWT middleware
3. Try using difference database

