import { writable } from 'svelte/store';

function createAuthStore() {
	// Initialize with empty values
	const auth = writable({
		token: '',
		role: ''
	});

	// Load from localStorage only in the browser
	if (typeof window !== 'undefined') {
		const storedToken = localStorage.getItem('token') || '';
		const storedRole = localStorage.getItem('role') || '';
		auth.set({ token: storedToken, role: storedRole });
	}

	// Subscribe to store and persist changes to localStorage
	auth.subscribe(({ token, role }) => {
		if (typeof window !== 'undefined') {
			if (token) {
				localStorage.setItem('token', token);
				localStorage.setItem('role', role);
			} else {
				localStorage.removeItem('token');
				localStorage.removeItem('role');
			}
		}
	});

	return auth;
}

// Create and export the store
export const authStore = createAuthStore();

// Logout function
export function logout() {
	authStore.set({ token: '', role: '' });
}
