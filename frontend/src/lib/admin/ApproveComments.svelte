<script>
	import { run } from 'svelte/legacy';

	import { onMount } from 'svelte';
	let comments = $state([]);
	run(() => {
		console.log(comments);
	});

	async function getComments() {
		try {
			const response = await fetch(`${import.meta.env.VITE_API_URL}/get-unapproved-comments`);

			if (!response.ok) {
				throw new Error('Failed to fetch unapproved comments');
			}

			const data = await response.json();
			comments = data;
		} catch (error) {
			console.error('Error fetching unapproved comments:', error);
		}
	}

	async function approveComment(id) {
		console.log('approveComment', id);
	}

	async function deleteComment(id) {
		console.log('deleteComment', id);
	}

	onMount(async () => {
		await getComments();
	});
</script>

<div class="container mx-auto p-6">
	<h1 class="text-2xl font-bold mb-6">Unapproved Comments</h1>
	{#if comments.length > 0}
		<div>
			{#each comments as comment}<div
					class="card p-4 shadow-md border border-surface-300-600-token"
				>
					<h2 class="text-xl font-semibold">{comment.title}</h2>
					<p class="text-sm text-surface-600-300-token">{comment.sub_title}</p>
				</div>
				<div class="mt-2 text-sm text-surface-600-300-token">
					<b>{comment.name}</b> - {new Date(comment.created_at * 1000).toLocaleDateString('en-US', {
						year: 'numeric',
						month: 'long',
						day: 'numeric'
					})}
				</div>

				<p class="mt-2 text-md">{comment.content}</p>

				<div class="flex gap-4 mt-4">
					<button class="btn variant-filled-primary" onclick={() => approveComment(comment.ID)}>
						✅ Approve
					</button>
					<button class="btn variant-filled-error" onclick={() => deleteComment(comment.ID)}>
						🗑 Delete
					</button>
				</div>
			{/each}
		</div>
	{:else}
		<div class="card p-4 shadow-md border border-surface-300-600-token">
			<p class="text-gray-500">No unapproved comments found.</p>
		</div>
	{/if}
</div>
