<script>
	import AllView from '$lib/blog/AllView.svelte';
	import BlogDetail from '$lib/blog/BlogDetail.svelte';

	let selectedBlogId = null;

	function handleBlogSelect(event) {
		selectedBlogId = event.detail;
	}
	function handleGoBackToAllBlogs() {
		selectedBlogId = null;
	}
</script>

<svelte:head>
	<title>Blog</title>
</svelte:head>

{#if selectedBlogId}
	<BlogDetail on:goBackToAllBlogs={handleGoBackToAllBlogs} blogId={selectedBlogId} />
{:else}
	<div class="min-h-screen">
		<AllView on:selectBlog={handleBlogSelect} />
	</div>
{/if}
