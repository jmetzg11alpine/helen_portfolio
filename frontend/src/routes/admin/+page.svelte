<script>
	import { authStore } from '$lib/admin/adminStore.js';
	import CreateBlogPost from '$lib/admin/CreateBlogPost.svelte';
	import ApproveComments from '$lib/admin/ApproveComments.svelte';
	import LogIn from '$lib/admin/LogIn.svelte';

	console.log($authStore);
</script>

<svelte:head>
	<title>Blog Admin</title>
</svelte:head>

<div class="min-h-screen w-full max-w-4xl mx-auto">
	{#if $authStore.role !== 'admin'}
		<LogIn />
	{:else}
		<CreateBlogPost />
		<ApproveComments />
	{/if}
</div>
