# ShopifyChallenge2021
[live app](https://shopify2021-57nddaoela-ue.a.run.app)

Feature completion
- [x] user registration (backend + frontend)
- [x] user login (backend + frontend)
- [x] graphql integration (frontend + backend)
- [x] password storage, hashing
- [x] file upload page
- [x] upload file to gcp bucket
- [x] update db: add file and append to user.files relationship
- [x] dashboard page displaying your images
- [x] display image tags in the upload screen
- [x] basic unit tests
- [x] integration tests, and decent unit test coverage
- [x] add docker-compose including test db image
- [x] CI/CD
- [ ] refactor save_file and get_file handlers to eliminate duplicate code
- [ ] server side input validation
- [ ] client side input validation
- [ ] unit test input validation
- [ ] general cleanup
- [ ] fix homepage flex display weirdness on smaller screens
- [ ] make an overview page
- [ ] make a demos page

## Run unit and integration tests
```bash
docker-compose -f docker-compose-test.yaml up --build
```
### Deploying to Prod
Currently deploy through CI only. This is triggered by merging a PR into master

## Graphql 
On dev mode you can GET at `/graphql` to use graphql playground

## Run app locally *unavailable for public use
* I could easily make a docker-compose for this (ie. to handle the dependencies), but the app relies on being able to access my gcloud account through a google application credentials file which would be an issue.

## Migrate (for local db)
### Migrate db
```bash
cd backend
go run main.go --migrate
```
### Migrate and serve db
```bash
cd backend
go run main.go --migrate-serve
```
### Nuke and recreate db
```bash
cd backend
go run main.go --nuke
```
