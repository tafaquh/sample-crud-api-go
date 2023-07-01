# Golang Simple CRUD
disclaimer: this is for learning purpose only. please do not make this repository structure as reference for your real project.

## How to run
```bash
$ go run main.go
```

## Setup
Please setup MYSQL database on your local machine and put the credentials into the .env file.
After you finish setup the MYSQL database, please open the mysql
```bash
~/go/src/github.com/tafaquh/coba-go                                                                                                                 11:29:54
> mysql -h 127.0.0.1 -uroot -p

```

and then do the followings command and costumize based your own needs
```bash
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

## Create Course
```bash
curl --request POST \
  --url http://localhost:1234/api/v1/course \
  --header 'Authorization: Basic eGMxUjJNdHJqTHhFdmRNbm9RbGZyWUVURGlGY0lQT0I6Ym80SlN1aW9RMVFmT3h3bzRNaFBGZXBPcWRCdUFhd2dKNklOZmpFdG9KWDZydUd5YUZUQ2NvZjluYTRJWEFXZw==' \
  --header 'Content-Type: application/json' \
  --data '{
	"company_id": "12",
	"name": "PT Cinta Sejati"
}'
```

## Get Course
```bash
curl --request GET \
  --url http://localhost:1234/api/v1/course/c9ddac4a-1396-4875-8f6a-6d0a64003a1f \
  --header 'Authorization: Basic eGMxUjJNdHJqTHhFdmRNbm9RbGZyWUVURGlGY0lQT0I6Ym80SlN1aW9RMVFmT3h3bzRNaFBGZXBPcWRCdUFhd2dKNklOZmpFdG9KWDZydUd5YUZUQ2NvZjluYTRJWEFXZw=='
```

## Register
```bash
curl --request POST \
  --url http://localhost:1234/api/v1/user/register \
  --header 'Authorization: Basic eGMxUjJNdHJqTHhFdmRNbm9RbGZyWUVURGlGY0lQT0I6Ym80SlN1aW9RMVFmT3h3bzRNaFBGZXBPcWRCdUFhd2dKNklOZmpFdG9KWDZydUd5YUZUQ2NvZjluYTRJWEFXZw==' \
  --header 'Content-Type: application/json' \
  --data '{
	"first_name": "test",
	"last_name": "register",
	"email": "test@register.com",
	"password": "123",
	"role": "student",
	"image": "https://images.pexels.com/photos/771742/pexels-photo-771742.jpeg?cs=srgb&dl=pexels-mohamed-abdelghaffar-771742.jpg&fm=jpg",
	"nip": "12345"
}'
```

## Login
```bash
curl --request POST \
  --url http://localhost:1234/api/v1/user/login \
  --header 'Authorization: Basic eGMxUjJNdHJqTHhFdmRNbm9RbGZyWUVURGlGY0lQT0I6Ym80SlN1aW9RMVFmT3h3bzRNaFBGZXBPcWRCdUFhd2dKNklOZmpFdG9KWDZydUd5YUZUQ2NvZjluYTRJWEFXZw==' \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "test@register.com",
	"password": "123"
}'
```