import { writable } from "svelte/store";

const stored_email = localStorage.getItem("email");
export const email = writable(stored_email);

email.subscribe(value => {
    localStorage.setItem('email', value === 'null' ? '' : value);
});
