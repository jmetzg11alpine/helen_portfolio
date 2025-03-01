import adapter from '@sveltejs/adapter-node';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import dotenv from 'dotenv';

dotenv.config();

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte'],
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter(),
		csrf: { checkOrigin: false },
		prerender: {
			entries: ['*']
		}
	}
};

export default config;
