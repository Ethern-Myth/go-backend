# GO - Simple Backend with SQLite

Simple Go REST API with SQLite database. 
Users tables containing Name, Surname, Email and Password.

---
#### Begin with :

```
go get
```
Then run
```
go run main.go
```

Application will be initialized and database will be created.
Open http://localhost:5000/user

#### Routes
---
* GET       /user
* GET       /user/:id&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;example /user/1
* POST      /user           
* PUT       /user/:id&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;example /user/1
* DELETE    /user/:id&nbsp;&nbsp;&nbsp;example /user/1

---

**Complex REST API will be coming soon with authentication**

**Vote which database to try out**:

MySQL, PostgreSQL or stay with SQLite
