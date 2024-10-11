import { writable } from 'svelte/store';

// Initialize users store with data from local storage or an empty array
const initialUsers = JSON.parse(localStorage.getItem('users')) || [];
export const users = writable(initialUsers);

// Initialize currentUser store with data from local storage or null
const initialCurrentUser = JSON.parse(localStorage.getItem('currentUser')) || null;
export const currentUser = writable(initialCurrentUser);

// Update local storage whenever users store changes
users.subscribe(value => {
  localStorage.setItem('users', JSON.stringify(value));
});

// Update local storage whenever currentUser store changes
currentUser.subscribe(value => {
  localStorage.setItem('currentUser', JSON.stringify(value));
});

export const currentAccount = writable(null);