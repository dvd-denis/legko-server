# legko-server
#### **REST API SEVER**
- > get all groups ->
`/groups`
- > get group articles or search group articles  ->
`/articles/:group_id?search=string`
- > get article steps->
`/steps/:article_id`
- > [**POST**] delete group ->
`/group/delete/:id`
- > [**POST**] create group ->
`/group`
- > [**POST**] create article ->
`/article`
- > [**POST**] create step ->
`/step`
- > [**POST**] create image ->
`/image`

#### In hosted on Heroku

#### Use postgresql 
- > up migrations ->
`migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" up`
- > drop migrations ->
`migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" down`
- > add new migration ->
`migrate create -ext sql -dir migrations create_articles`