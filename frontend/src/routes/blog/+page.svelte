<script>
	import AllView from '$lib/blog/AllView.svelte';
	import BlogDetail from '$lib/blog/BlogDetail.svelte';

	let selectedBlogId = $state(null);

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
	<BlogDetail on:goBackToAllBlogs={handleGoBackToAllBlogs} blogID={selectedBlogId} />
{:else}
	<div class="min-h-screen">
		<AllView on:selectBlog={handleBlogSelect} />
	</div>
{/if}
