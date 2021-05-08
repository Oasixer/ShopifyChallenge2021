# Architecture
**This document describes the high-level architecture of this project**

If you want to familiarize yourself with the code base and *generally* how it works, this is a good place to be.

## High Level TLDR

The frontend code gets loaded into `frontend/public/index`. `frontend/src/App.svelte` is the main svelte app. Everything in `frontend/src/pages/` is the routes. 

## Code Map

#### Code Map Legend

`<file name>` for a file name

`<folder name>/` for a folder

`<folder name>/<file name>` for a file within a folder

### `frontend/public`

This is all the publicly accessible resources to be served. This includes(but not limited to) images, css, html, and javascript.

### `frontend/package.json`

This is the package list for javascript. It also includes workspaces and specialty commands.

### `frontend/webpack.config.js`

This is the config for webpack. Webpack is the bundling system we use for svelte, css, and html.

### `frontend/index.html`

This is the base html file where all of our svelte gets loaded.

### `frontend/public/build`

This is all the built svelte files. As svelte is a compiler, it needs to be built to javascript. This is where the files end up after being built.

### `frontend/src/`

This is the place where all of our svelte source code lives.

### `frontend/src/main.js`

This is the actual script that runs `frontend/src/App.svelte`. It is also the thing that gets loaded by `frontend/public/index.html`.

### `frontend/src/App.svelte`

This is the core and start of our svelte project. All code gets loaded and runs from here.

### `frontend/src/pages`

This is where the frontend routes live. We use [Routify](https://routify.dev/) for our frontend router so this file structure is based off of this. You can read more about it [here](https://routify.dev/guide/introduction/structure).

### `frontend/src/components`

This is where our components are put.
