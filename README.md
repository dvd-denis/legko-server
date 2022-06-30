# legko-server
#### **REST API SEVER**
- > hello ->
`/hello`
- > get alls articles ->
`/articles`

#### Use postgresql 
- > up migrations ->
`migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" up`
- > drop migrations ->
`migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" down`
- > add new migration ->
`migrate create -ext sql -dir migrations create_articles`