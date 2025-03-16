<script>
	import { createEventDispatcher, onMount } from 'svelte';
	import { getModalStore } from '@skeletonlabs/skeleton';
	let { blogID } = $props();
	const dispatch = createEventDispatcher();
	let blog = $state({});

	onMount(async () => {
		try {
			const url = `${import.meta.env.VITE_API_URL}/blog/${blogID}`;
			const response = await fetch(url);

			if (!response.ok) {
				throw new Error('Failed to fetch blog content');
			}

			const data = await response.json();
			blog = data;
		} catch (error) {
			console.error('Error fetching blog content:', error);
		}
	});

	function handleGoBack() {
		dispatch('goBackToAllBlogs');
	}
	const modalStore = getModalStore();
	const modal = {
		type: 'component',
		component: 'commentModal',
		meta: { blogID }
	};
	function openCommentModal() {
		modalStore.trigger(modal);
	}
</script>

<svelte:head>
	<title>Blog {blogID}</title>
</svelte:head>

<div class="container min-h-screen mx-auto px-4 max-w-3xl">
	<div class="flex justify-between mb-8">
		<button onclick={handleGoBack} class="btn variant-ghost-tertiary"
			><i class="fas fa-arrow-left mr-2"></i> Go Back
		</button>
		<button class="btn variant-ghost-secondary" onclick={openCommentModal}> Leave Comment </button>
	</div>
	<article class="prose prose-lg dark:prose-invert">
		<div class="flex justify-between items-center mb-2">
			<h1 class="text-4xl font-bold m-0">{blog.title}</h1>
			<div class="text-sm text-surface-600-300-token">
				{new Date(blog.created_at * 1000).toLocaleDateString('en-US', {
					year: 'numeric',
					month: 'long',
					day: 'numeric'
				})}
			</div>
		</div>
		<h2 class="text-xl text-surface-600-300-token mb-4">{blog.sub_title}</h2>
		<div class="blog-content leading-relaxed mb-12">{blog.content}</div>
	</article>

	<section class="mt-16 border-t border-surface-300-600-token pt-8">
		<h3 class="text-2xl font-bold mb-8">Comments</h3>
		{#if blog.Comments && blog.Comments.length > 0}
			<div class="space-y-8">
				{#each blog.Comments as comment}
					<div class="card p-4 variant-soft">
						<h4 class="font-semibold mb-2">{comment.name}</h4>
						<div class="mb-3">{comment.content}</div>
						<div class="text-sm text-surface-600-300-token">
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
			<div class="card p-4 variant-soft text-center">No comments yet</div>
		{/if}
	</section>
</div>

<style>
	.blog-content {
		white-space: pre-wrap;
	}
</style>
