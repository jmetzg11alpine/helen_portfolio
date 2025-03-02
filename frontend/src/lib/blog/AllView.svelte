<script>
	import { onMount } from 'svelte';
	let blogPosts = [];
	onMount(async () => {
		try {
			const response = await fetch(`${import.meta.env.VITE_API_URL}/blog-previews`);
			if (!response.ok) {
				throw new Error('Failed to fetch blog posts');
			}
			const data = await response.json();
			blogPosts = data;
		} catch (error) {
			console.error('Error fetching blog posts:', error);
		}
	});
</script>

<div class="container mx-auto p-4">
	<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
		{#each blogPosts as post}
			<div class="card p-4 variant-tertiary-primary">
				<header class="card-header">
					<h3 class="h3 font-bold">{post.title}</h3>
				</header>
				<section class="p-4">
					<p class="opacity-80">{post['sub_title']}</p>
					<p class="text-sm opacity-60 mt-2">
						{new Date(post.created_at * 1000).toLocaleDateString('en-US', {
							year: 'numeric',
							month: 'short',
							day: 'numeric'
						})}
					</p>
				</section>
				<footer class="card-footer">
					<button class="btn variant-filled variant-filled-secondary">Read More</button>
				</footer>
			</div>
		{/each}
	</div>
</div>
