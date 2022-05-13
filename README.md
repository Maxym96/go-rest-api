#rest-api

#user-service

GET: /users -- list of users -- 200, 404, 500
GET: /user/:id -- get user by id -- 200, 404, 500
POST: /user/:id -- create user by id -- 204, 4xx, Header Location: url, 500
PUT: /user/:id -- fully update user by id -- 200/204, 404, 500
PATCH: /user/:id -- partially update user by id -- 200/204, 404, 500
DELETE: /user/:id -- delete user by id -- 200/204, 404, 500