<script>
	import { fade } from 'svelte/transition';
</script>

<header class="w-full max-w-6xl flex justify-between items-center p-4" transition:fade>
	<h1 class="text-2xl font-bold text-primary-500">Helen Metzger</h1>
	<nav class="hidden md:flex">
		<ul class="flex gap-x-6">
			<li><a href="#about" class="btn-link">About</a></li>
			<li><a href="#books" class="btn-link">Books</a></li>
			<li><a href="#newsletter" class="btn-link">News Letter</a></li>
			<li><a href="#contact" class="btn-link">Contact</a></li>
		</ul>
	</nav>
</header>

<style>
	.btn-link {
		@apply text-primary-500 hover:underline focus:outline-none focus:ring-2 focus:ring-primary-500;
	}
</style>
