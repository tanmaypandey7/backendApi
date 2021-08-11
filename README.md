Get started
===========


Usage
-----

### 1. Create Operations

Inserts the product details into the database and returns its id.

#### HTTP Request
```bash
POST http://127.0.0.1:8080/api/product
```


Query Params:

| Sr. No      | Param       | Type          |  
| :---        |    :----:   |          ---: |  
| 1.      | Name                    | str   | yes | 
| 2.      | DOB              | str   | yes | 
| 3.      | Address       | str   | yes | 
| 4.      | Description       | str   | yes | 


#### Response
If successful, this method returns a response like:
```bash
{
    "data": {
        "InsertedID": "6113ed484ea2632dc0b61605"
    },
    "message": "Record inserted successfully",
    "success": true
}
```
#### Errors
The following table identifies error messages that the API could return in response to a call to this method.

| Error Type      | Error Detail       |
| :---        |    :----:   |          
| Record failed to insert(500)     | Service was unable to insert record in DB|    

### 2. Read Operations

Returns the details of product which matches the id.

#### HTTP Request
```bash
GET http://127.0.0.1:2020/api/records/id
```

#### Response
If successful, this method returns a response body like:
```bash
{
    "data": {
        "record": {
            "Name": "abc",
            "DOB": "11/01/1991",
            "Address": "def",
            "Description": "xyz",
            "CreatedAt": "2021-08-11T15:31:20.911Z"
        }
    },
    "success": true
}
```

#### Errors
The following table identifies error messages that the API could return in response to a call to this method.

| Error Type      | Error Detail       |
| :---        |    :----:   |          
| BadRequest(400)     | Server couldn't parse the ID |    
| NotFound(404)      | The ID doesn't exist              | 

### 3. Update Operations

Updates the database according to the given values

#### HTTP Request
```bash
PUT http://127.0.0.1:2020/api/records/id
```

Query Params:

| Sr. No      | Param       | Type          |  
| :---        |    :----:   |          ---: |  
| 1.      | Name                    | str   | yes | 
| 2.      | DOB              | str   | yes | 
| 3.      | Address       | str   | yes | 
| 4.      | Description       | str   | yes | 


#### Response
If successful, this method returns a response like:
```bash
{
    "data": {
        "record": {
            "Name": "abcccsz",
            "DOB": "11/01/1990",
            "Address": "def",
            "Description": "xyz",
            "CreatedAt": "2021-08-11T15:17:40.014Z"
        }
    },
    "success": true
}
```
#### Errors

| Error Type      | Error Detail       |
| :---        |    :----:   |          
| BadRequest(400)     | Server couldn't parse the ID or JSON |    
| DoesNotExist(400)     | Product with given name already exists|    
| NotFound(404)      | The ID doesn't exist              | 

### 4. Delete Operations

Deletes a value in the database given its id.

#### HTTP Request 
```bash
DELETE http://127.0.0.1:2020/api/records/id
```


#### Response
If successful, this method returns a response like:
```bash
{
    "message": "Record deleted successfully",
    "success": true
}
```


#### Errors
The following table identifies error messages that the API could return in response to a call to this method.

| Error Type      | Error Detail       |
| :---        |    :----:   |          
| BadRequest(400)     | Server couldn't parse the ID or JSON |    
| NotFound(404)      | The ID doesn't exist              | 





License
=======

Apache License 2.0. For more details, please read the
[LICENSE](https://github.com/imWildCat/scylla/blob/master/LICENSE) file.
