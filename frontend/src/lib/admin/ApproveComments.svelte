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
		const url = `${import.meta.env.VITE_API_URL}/approve-comment`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id })
		});

		if (!response.ok) {
			throw new Error('Failed to approve comment');
		}

		await getComments();
	}

	async function deleteComment(id) {
		const url = `${import.meta.env.VITE_API_URL}/delete-comment`;
		const response = await fetch(url, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id })
		});

		if (!response.ok) {
			throw new Error('Failed to delete comment');
		}

		await getComments();
	}

	onMount(async () => {
		await getComments();
	});
</script>

<div class="mx-auto">
	<h1 class="text-2xl font-bold mb-6">Unapproved Comments</h1>
	{#if comments.length > 0}
		<div>
			{#each comments as comment}
				<div class="card p-4 shadow-md bg-surface-100 border border-surface-400 mb-4">
					<h2 class="text-xl font-semibold">{comment.title}</h2>

					<div class="mt-2 text-sm text-surface-700-300">
						<b>{comment.name}</b> - {new Date(comment.created_at * 1000).toLocaleDateString(
							'en-US',
							{
								year: 'numeric',
								month: 'long',
								day: 'numeric'
							}
						)}
					</div>

					<p class="mt-3 text-md">{comment.content}</p>

					<div class="flex justify-between mt-4">
						<button
							class="btn preset-filled-success-500"
							onclick={() => approveComment(comment.ID)}
						>
							✅ Approve
						</button>
						<button class="btn preset-filled-error-500" onclick={() => deleteComment(comment.ID)}>
							🗑 Delete
						</button>
					</div>
				</div>
			{/each}
		</div>
	{:else}
		<div class="card p-4 shadow-md border border-surface-300-700">
			<p class="text-primary-500">No unapproved comments found.</p>
		</div>
	{/if}
</div>
