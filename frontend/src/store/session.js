import { writable } from "svelte/store";

const stored_sessionToken = localStorage.getItem("sessionToken");
export const sessionToken = writable(stored_sessionToken);

sessionToken.subscribe(value => {
    localStorage.setItem('sessionToken', value === 'null' ? '' : value);
});
