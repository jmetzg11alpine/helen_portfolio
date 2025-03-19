<script>
	import AllView from '$lib/blog/AllView.svelte';
	import BlogDetail from '$lib/blog/BlogDetail.svelte';
	import CommentModal from '$lib/blog/CommentModal.svelte';

	import { selectedBlogId } from '$lib/blog/blogStore.js';
</script>

<svelte:head>
	<title>Blog</title>
</svelte:head>

{#if $selectedBlogId}
	<BlogDetail />
{:else}
	<div>
		<AllView />
	</div>
{/if}

<CommentModal />
