# legko-server
#### **REST API SEVER**
- > get all articles ->
`/articles`
- > get steps from article ->
`/step/:id`
- > delete article ->
`/article/delete/:id`
- > get questions from article ->
`/questions/:id`
- > [**POST**] create article ->
`/article`
- > [**POST**] create step ->
`/step`
- > [**POST**] create image ->
`/image`


#### Use postgresql 
- > up migrations ->
`migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" up`
- > drop migrations ->
`migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" down`
- > add new migration ->
`migrate create -ext sql -dir migrations create_articles`