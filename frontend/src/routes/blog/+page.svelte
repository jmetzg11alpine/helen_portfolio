<script>
	import AllView from '$lib/blog/AllView.svelte';
	import BlogDetail from '$lib/blog/BlogDetail.svelte';

	let selectedBlogId = null;

	function handleBlogSelect(event) {
		selectedBlogId = event.detail;
		console.log(selectedBlogId);
	}
	function handleGoBackToAllBlogs() {
		selectedBlogId = null;
	}
</script>

<svelte:head>
	<title>Blog</title>
</svelte:head>

{#if selectedBlogId}
	<h1>Blog Detail</h1>
	<BlogDetail on:goBackToAllBlogs={handleGoBackToAllBlogs} blogID={selectedBlogId} />
{:else}
	<h1>All Blogs</h1>
	<div class="min-h-screen">
		<AllView on:selectBlog={handleBlogSelect} />
	</div>
{/if}
