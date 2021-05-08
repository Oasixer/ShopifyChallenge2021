import { writable } from "svelte/store";

const stored_email = localStorage.getItem("email");
export const email = writable(stored_email);

email.subscribe(value => {
    localStorage.setItem('email', value === 'null' ? '' : value);
});

const stored_username = localStorage.getItem("username");
export const username = writable(stored_username);

username.subscribe(value => {
    localStorage.setItem('username', value === 'null' ? '' : value);
});
