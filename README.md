# ShopifyChallenge2021
## UPDATE MAY 21
I fixed all the problems / incomplete features with this app (after the deadline), see ['v2' branch instead](htts)

## Old README
~~Unfortunately, I let the scope of this project creep too wide to wrap it up in a satisfactory way by the deadline. However, you can see a live, semi-working demo of what I have, here~~
https://shopify2021-57nddaoela-ue.a.run.app

Feature completion
- [x] user registration (backend + frontend)
- [x] user login (backend + frontend)
- [x] graphql integration (frontend + backend)
- [x] password storage, hashing
- [x] file upload page
- [x] upload file to gcp bucket
- [ ] update db: add file and append to user.files relationship !!!this is what I got stuck on for a day - for some reason I can't get my GORM to properly update this relationship, no matter what I tried
- [ ] dashboard page displaying your images (BLOCKED by the above - dashboard page currently just )
- [x] display image tags in the upload screen with suggestions, autofill, etc. [ this is semi complete but the code is spaghetti because I was trying to get the tag logic fixed. ]

- [x] basic unit tests
- [ ] integration tests, and decent unit test coverage
- [ ] add docker-compose including an image for a local db so that shopify team can run the app w/o needing to create local postgres instance

## Run app locally
* TODO: Once docker-compose is working for local development / testing, put command here

## Run backend (must have built frontend locally, see frontend/README.md)
```bash
cd backend
go run main.go
```

### Deploying to Prod
We currently deploy through CI only. This is triggered by merging a PR into the `prod` branch.

## Graphql 

On dev mode you can GET at `/graphql` to get graphql playground

## Architecture

This can be read at `ARCHITECTURE.md`

## Migrate
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
