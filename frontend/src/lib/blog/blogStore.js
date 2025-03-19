import { writable } from 'svelte/store';

export const selectedBlogId = writable(null);

export const modalOpen = writable(false);
