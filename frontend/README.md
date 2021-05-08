# FRONTEND

**ALL THESE INSTRUCTIONS ARE TO BE USED INSIDE `<MAIN>/frontend`**

## Install
### Required dependencies
To install all *required* frontend dependencies, run `yarn install`

### Optional dependencies
* Optional dev features hot reloading and sound cues rely on the dependencies listed below which are currently only tested on linux
	* if you are onboarding to the project and use a different OS, let us know and we'll help you port the install instructions to your OS as needed

#### Inotify-Tools
Required for hot reloading

##### Arch
[install inotify-tools from aur](https://archlinux.org/packages/community/x86_64/inotify-tools/)

##### Debian/Ubuntu
`apt install -y inotify-tools`

#### Pulse Audio
Required for hot reloading with fancy sound cues when builds fail/succeed

##### Arch
[install pulseaudio from aur](https://archlinux.org/packages/extra/x86_64/pulseaudio/)

##### Debian/Ubuntu
`apt install -y pulseaudio`

## Build

To build svelte to be served, run `yarn build`

### With hot reloading
To rebuild svelte every time a file is changed in the frontend/src subtree, run `yarn run dev`

#### With sound cues
Hot reloading like above, but with a generic error/success sound so that you know when a build failed without having the build window open `yarn run devSound`
* Requires pulseaudio (see above)

## Architecture

This can be read at `ARCHITECTURE.md`
