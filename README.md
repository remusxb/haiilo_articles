# Haiilo Articles

## Running the application
From within the project root run `docker-compose up --build` to build and run the app. 
The `--build` tag is only needed when running the application for the  first time.

The database will be created & populated as specified in `db/postgres/structure.sql` & `/db/postgres/data.sql`; 
These will be mounted in `docker-entrypoint-initdb.d` which will automatically create/populate tables 
when app starts for the first time.
The  username/password/dbname are specified in the env from the `db` service.

### Sending requests

Easiest way is to import into `Postman` the `Haiilo.postman_collection.json` file.

OR

* GET `/articles` - list articles
  * `curl -v --location --request GET 'http://127.0.0.1:8080/articles'`

* POST `/articles` - create article
  * `curl --location --request POST "http://0.0.0.0:8080/articles" \
    --header "Content-Type: application/json" \
    --data-raw "{
    \"title\": \"article 1\",
    \"link\": \"/article-1\"
    }"` - this creates an article with `title: "article 1"` and `link: "/article-1"`

## Running tests
**NOTE: MUST BE RUN FROM THE `/server` FOLDER(run `cd server/` from the root folder)**
* `go test -v -tags=unit ./internal/articles/ `

Yes... there could've been integration tests also, alongside a function/package that creates a container specifically for tests
(so we don't pollute the main db + the tests would fail 2nd time due to duplicates), with dedicated db. But I would've also  needed to add a migration option.
