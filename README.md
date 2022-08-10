# GIN-API

# Api Routes

1. GET /teapot
2. POST /file
3. POST /file/chunk
4. POST /file/chunk/merge
5. GET /file/chunk/state
6. POST /auth/signup
7. POST /auth/login
8. GET /auth/me
9. POST /s3/r2/upload
10. POST /image/classify/file
11. GET /posts/:id
12. DELETE /posts/:id
13. POST /posts
14. GET /posts
15. GET /projects
16. POST /projects
17. GET /projects/:id
18. DELETE /projects/:id
19. GET /healthz
20. GET /

# What are this project build with

1. gin for web framework
2. gorm for orm 
3. redis for cache
4. grpc for rpc
5. jwt for auth
6. godotenv for config
7. docker for deploy
8. s3 like for storage

# What did this project do

1. basic auth build on jwt
2. big file chunk upload
3. CRUD and cache posts
4. CRUD and cache projects
5. Image classify (nsfw) build on grpc
6. upload limit: size and mime type (middleware)
7. rate limit
