import { writable } from "svelte/store";

const stored_fileUpload = localStorage.getItem("fileUpload");
export const fileUpload = writable(stored_fileUpload);

fileUpload.subscribe(value => {
    localStorage.setItem('fileUpload', value === 'null' ? '' : value);
});
