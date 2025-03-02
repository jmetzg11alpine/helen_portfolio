<script>
	import { fade } from 'svelte/transition';
	import { page } from '$app/stores';

	function handleNavigation(event, href) {
		// If we're not on the home page and it's a hash link
		if ($page.url.pathname !== '/' && href.startsWith('#')) {
			event.preventDefault();
			window.location.href = '/' + href;
		}
	}
</script>

<header class="w-full max-w-6xl flex justify-between items-center p-4" transition:fade>
	<h1 class="text-2xl font-bold text-primary-500">Helen Metzger</h1>
	<nav class="hidden md:flex">
		<ul class="flex gap-x-6">
			<li>
				<a href="#about" class="btn-link" on:click={(e) => handleNavigation(e, '#about')}>About</a>
			</li>
			<li>
				<a href="#books" class="btn-link" on:click={(e) => handleNavigation(e, '#books')}>Books</a>
			</li>
			<li>
				<a href="#newsletter" class="btn-link" on:click={(e) => handleNavigation(e, '#newsletter')}
					>News Letter</a
				>
			</li>
			<li>
				<a href="/blog" class="btn-link" on:click={(e) => handleNavigation(e, '/blog')}>Blog</a>
			</li>
			<li>
				<a href="/merch" class="btn-link" on:click={(e) => handleNavigation(e, '/merch')}>Merch</a>
			</li>
			<li>
				<a href="#contact" class="btn-link" on:click={(e) => handleNavigation(e, '#contact')}
					>Contact</a
				>
			</li>
		</ul>
	</nav>
</header>

<style>
	.btn-link {
		@apply text-primary-500 hover:underline focus:outline-none focus:ring-2 focus:ring-primary-500;
	}
</style>
