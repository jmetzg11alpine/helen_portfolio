<script>
	import CreateBlogPost from '$lib/admin/CreateBlogPost.svelte';
	import ApproveComments from '$lib/admin/ApproveComments.svelte';
</script>

<svelte:head>
	<title>Blog Admin</title>
</svelte:head>

<div class="min-h-screen w-full max-w-4xl mx-auto">
	<CreateBlogPost />
	<ApproveComments />
</div>
