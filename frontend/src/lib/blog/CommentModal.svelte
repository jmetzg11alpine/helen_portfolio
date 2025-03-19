<script>
	import { Modal } from '@skeletonlabs/skeleton-svelte';
	import { modalOpen, selectedBlogId } from './blogStore.js';

	// Use $state for reactive variables in Svelte 5
	let nameInput = $state('');
	let commentInput = $state('');
	let submitError = $state('');
	let submitSuccess = $state(false);

	async function submitComment() {
		if (!nameInput || !commentInput) {
			submitError = 'Please fill out both fields';
			return;
		}

		try {
			const response = await fetch(`${import.meta.env.VITE_API_URL}/blog-comment`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					name: nameInput,
					comment: commentInput,
					blogID: $selectedBlogId
				})
			});

			if (!response.ok) {
				throw new Error('Failed to submit comment');
			}

			// Success handling
			submitSuccess = true;
			submitError = '';

			// Reset form after successful submission
			setTimeout(() => {
				nameInput = '';
				commentInput = '';
				modalOpen.set(false);
				submitSuccess = false;
			}, 1500);
		} catch (error) {
			console.error('Error submitting comment:', error);
			submitError = 'Error submitting comment. Please try again.';
		}
	}
</script>

<Modal
	open={$modalOpen}
	onOpenChange={(e) => modalOpen.set(e.detail.open)}
	triggerBase="btn preset-tonal"
	contentBase="card bg-surface-100-900 p-4 space-y-4 shadow-xl max-w-screen-sm"
	backdropClasses="backdrop-blur-sm"
>
	{#snippet content()}
		<header class="flex justify-between items-center mb-4">
			<h2 class="h2 text-tertiary-500">Leave a Comment</h2>
		</header>

		{#if submitSuccess}
			<div class="alert variant-filled-success p-4 mb-4">Comment submitted successfully!</div>
		{:else}
			{#if submitError}
				<div class="alert variant-filled-error p-4 mb-4">
					{submitError}
				</div>
			{/if}

			<div class="space-y-4">
				<label class="label">
					<span>Your Name</span>
					<input
						type="text"
						bind:value={nameInput}
						placeholder="Your Name"
						class="input p-3 w-full border border-surface-300 dark:border-surface-700 rounded"
					/>
				</label>

				<label class="label">
					<span>Your Comment</span>
					<textarea
						bind:value={commentInput}
						placeholder="Share your thoughts..."
						class="textarea p-3 w-full h-32 border border-surface-300 dark:border-surface-700 rounded"
					></textarea>
				</label>
			</div>

			<footer class="flex justify-end gap-4 mt-6">
				<button
					type="button"
					class="btn variant-ghost-surface"
					onclick={() => modalOpen.set(false)}
				>
					Cancel
				</button>
				<button type="button" class="btn variant-filled-primary" onclick={submitComment}>
					Submit Comment
				</button>
			</footer>
		{/if}
	{/snippet}
</Modal>
