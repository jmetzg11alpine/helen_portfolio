<script>
	import '../app.postcss';
	import Header from '$lib/Header.svelte';
</script>

<div
	class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br variant-gradient-error-warning text-center p-6"
>
	<Header />
	<slot />
</div>
