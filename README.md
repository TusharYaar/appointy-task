# Details

Made for appointy task<br>
Made by Tushar S Agrawal<br>
Reg no: 19BCE0559<br>
Submitted on 9th October, 2021<br>

# Installation

clone the repo

make sure you are in root folder

run `go get go.mongodb.org/mongo-driver/mongo` to install mongo drivers

use `go run main.go` to start the server

#### For Testing

`cd testing`

`go test`

# Routes

- [x] `POST /user` adds new user<br>
- [x] `GET /user/<user_id>` returns details of specified user<br>
- [x] `POST /post` adds new post<br>
- [x] `GET /post/<post_id>` returns details of specified post<br>
- [x] `GET /post/user/<user_id>` returns all post by specific user

# Testing

`cd testing`

- [x] `/user` - 5 tests cases<br>
- [x] `/user/<user_id>` - 4 tests cases<br>
- [x] `/post` - 5 tests cases<br>
- [x] `/post/<post_id>` - 4 tests cases<br>
- [x] `/post/user/<user_id>`- 3 tests cases<br>

# About Database

1. Uses mongodb Atlas, and connects via url.<br>
2. uses a temporary user<br>

# About Data

4 accounts are available now.

```
id: 6161447d56188e12db944c80
email: tusharsagrawal6@gmail.com
name: Tushar S Agrawal
number of posts: 5


id: 6160b6d57da33296d38ef6a6
email: ema23il2@gmail.com
name: sdcvsdc
number of posts: 4

id: 6160baddd5efcdffb10eab63
email: emal2@gmail.com
name: sdvsdvc
number of posts: 1

id: 61619a2d51169d1757ad2021
name: "test"
email: "test@test.com"
number of posts: 1
```

# Folder Structure

```
|- root
    |-connection
        |-connection.go
    |-handlers                      //contains route handlers
        |-create_post.go            // POST /post
        |-create_user.go            // POST /user
        |-get_post.go               // GET /post/<post_id>
        |-get_posts_by_user.go      // GET /posts/user/<user_id>
        |-get_user.go               // GET /user/<user_id>
    |-models                        //contains models for user and post
        |-post.go                   //post struct
        |-user.go                   //user struct
    |-testing                       //contains unit test for handlers
        |-create_post_test.go       //POST /post
        |-create_user_test.go       // POST /user
        |-get_post_test.go          // GET /post/<post_id>
        |get_post_by_user_test.go   //GET /posts/user/<user_id>
        |get_user_test.go           // GET /user/<user_id>
    |-go.mod
    |-go.sum
    |-main.go                       //Main Fnction
    |-README.md

```
