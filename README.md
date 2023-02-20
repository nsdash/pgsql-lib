# Coverage for pgsql

___

## Usage

* Create .env file and set params
```
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
DB_DRIVER=
```

* Use Manager
```go
manager := pgsql.NewSqlManager()

manager.Exec("...Query")
```
