<script>
	import { run } from 'svelte/legacy';

	import { fade } from 'svelte/transition';
	import { page } from '$app/stores';
	let isMenuOpen = $state(false);

	run(() => {
		if (typeof document !== 'undefined') {
			if (isMenuOpen) {
				document.body.style.overflow = 'hidden';
			} else {
				document.body.style.overflow = '';
			}
		}
	});

	function handleNavigation(event, href) {
		// If we're not on the home page and it's a hash link
		if ($page.url.pathname !== '/' && href.startsWith('#')) {
			event.preventDefault();
			window.location.href = '/' + href;
		}
		isMenuOpen = false;
	}

	function toggleMenu() {
		isMenuOpen = !isMenuOpen;
	}
</script>

<header class="w-full max-w-6xl flex justify-between items-center p-4" transition:fade>
	<h1 class="text-2xl font-bold text-success-500">Helen Metzger</h1>

	<button class="md:hidden z-50" onclick={toggleMenu} aria-label="Toggle menue">
		<div class="hamburger {isMenuOpen ? 'open' : ''}">
			<span></span>
			<span></span>
			<span></span>
		</div>
	</button>

	{#if isMenuOpen}
		<nav class="fixed inset-0 bg-white z-40 md:hidden h-screen overflow-hidden" transition:fade>
			<ul class="flex flex-col items-center justify-center h-full gap-y-6">
				<li>
					<a href="#about" class="btn-link text-xl" onclick={(e) => handleNavigation(e, '#about')}
						>About</a
					>
				</li>
				<li>
					<a href="#books" class="btn-link text-xl" onclick={(e) => handleNavigation(e, '#books')}
						>Books</a
					>
				</li>
				<li>
					<a
						href="#newsletter"
						class="btn-link text-xl"
						onclick={(e) => handleNavigation(e, '#newsletter')}>News Letter</a
					>
				</li>
				<li>
					<a href="/blog" class="btn-link text-xl" onclick={(e) => handleNavigation(e, '/blog')}
						>Blog</a
					>
				</li>
				<li>
					<a href="/merch" class="btn-link text-xl" onclick={(e) => handleNavigation(e, '/merch')}
						>Merch</a
					>
				</li>
				<li>
					<a
						href="#contact"
						class="btn-link text-xl"
						onclick={(e) => handleNavigation(e, '#contact')}>Contact</a
					>
				</li>
			</ul>
		</nav>
	{/if}

	<nav class="hidden md:flex">
		<ul class="flex gap-x-6">
			<li>
				<a href="#about" class="btn-link" onclick={(e) => handleNavigation(e, '#about')}>About</a>
			</li>
			<li>
				<a href="#books" class="btn-link" onclick={(e) => handleNavigation(e, '#books')}>Books</a>
			</li>
			<li>
				<a href="#newsletter" class="btn-link" onclick={(e) => handleNavigation(e, '#newsletter')}
					>News Letter</a
				>
			</li>
			<li>
				<a href="/blog" class="btn-link" onclick={(e) => handleNavigation(e, '/blog')}>Blog</a>
			</li>
			<li>
				<a href="/merch" class="btn-link" onclick={(e) => handleNavigation(e, '/merch')}>Merch</a>
			</li>
			<li>
				<a href="#contact" class="btn-link" onclick={(e) => handleNavigation(e, '#contact')}
					>Contact</a
				>
			</li>
		</ul>
	</nav>
</header>
