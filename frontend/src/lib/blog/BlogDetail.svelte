<script>
	let blog = $state({});

	import { selectedBlogId, modalOpen } from './blogStore.js';

	$effect(async () => {
		const id = $selectedBlogId; // Access store value reactively

		// Ensure `id` is a valid number before fetching
		if (!id || isNaN(id)) return;

		try {
			const url = `${import.meta.env.VITE_API_URL}/blog/${id}`;
			const response = await fetch(url);

			if (!response.ok) {
				throw new Error('Failed to fetch blog content');
			}

			blog = await response.json();
		} catch (error) {
			console.error('Error fetching blog content:', error);
			blog = null; // Reset on error
		}
	});

	function handleGoBack() {
		selectedBlogId.set(null);
	}

	function openCommentModal() {
		openModal();
	}
</script>

<svelte:head>
	<title>Blog {selectedBlogId}</title>
</svelte:head>

<div class="container mx-auto px-4 max-w-3xl">
	<div class="flex justify-between mb-8">
		<button onclick={handleGoBack} class="btn preset-filled-tertiary-500 shadow-xl"
			><i class="fas fa-arrow-left mr-2"></i>Go Back
		</button>
		<button class="btn preset-filled-tertiary-500 shadow-xl" onclick={() => modalOpen.set(true)}>
			Leave Comment
		</button>
	</div>
	<article class="prose prose-lg dark:prose-invert">
		<div class="flex flex-col items-center mb-2">
			<h1 class="text-4xl font-bold m-0">{blog.title}</h1>
			<h2 class="text-xl text-surface-700-300 mt-3">{blog.sub_title}</h2>
			<div class="text-sm text-surface-700 mb-4">
				{new Date(blog.created_at * 1000).toLocaleDateString('en-US', {
					year: 'numeric',
					month: 'long',
					day: 'numeric'
				})}
			</div>
		</div>

		<div class="blog-content leading-relaxed mb-12">{blog.content}</div>
	</article>

	<section class="mt-16 border-t border-surface-300-700 pt-8">
		<h3 class="text-2xl font-bold mb-8">Comments</h3>
		{#if blog.Comments && blog.Comments.length > 0}
			<div class="space-y-8">
				{#each blog.Comments as comment}
					<div class="card p-4 preset-tonal">
						<h4 class="font-semibold mb-2">{comment.name}</h4>
						<div class="mb-3">{comment.content}</div>
						<div class="text-sm text-surface-700-300">
							{new Date(comment.created_at * 1000).toLocaleDateString('en-US', {
								year: 'numeric',
								month: 'long',
								day: 'numeric'
							})}
						</div>
					</div>
				{/each}
			</div>
		{:else}
			<div class="card p-4 preset-tonal text-center">No comments yet</div>
		{/if}
	</section>
</div>

<style>
	.blog-content {
		white-space: pre-wrap;
	}
</style>
