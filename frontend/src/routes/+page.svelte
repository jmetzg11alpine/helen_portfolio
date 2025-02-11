<script>
	import Header from '../lib/Header.svelte';
	import Books from '../lib/Books.svelte';
	import About from '../lib/About.svelte';
	import Contact from '../lib/Contact.svelte';
	import NewsLetter from '../lib/NewsLetter.svelte';
	import Footer from '../lib/Footer.svelte';
</script>

<div class="min-h-screen flex flex-col items-center justify-center bg-tertiary-200 text-center p-6">
	<Header />
	<About />
	<Books />
	<NewsLetter />
	<Contact />
	<Footer />
</div>
