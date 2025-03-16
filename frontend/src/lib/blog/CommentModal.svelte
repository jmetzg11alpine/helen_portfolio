<script>
	import { getModalStore, getToastStore } from '@skeletonlabs/skeleton';
	const modalStore = getModalStore();
	const toastStore = getToastStore();
	let { parent } = $props();

	function closeModal() {
		modalStore.close();
	}

	async function submitComment(event) {
		event.preventDefault();

		const nameInput = document.getElementById('name');
		const commentInput = document.getElementById('comment');

		if (!nameInput.value.trim()) {
			toastStore.trigger({
				message: 'Add a name',
				background: 'variant-filled-error',
				hideDismiss: true,
				timeout: 2000
			});
			return;
		}

		if (!commentInput.value.trim()) {
			toastStore.trigger({
				message: 'Add a comment',
				background: 'variant-filled-error',
				hideDismiss: true,
				timeout: 2000
			});
			return;
		}

		try {
			const response = await fetch(`${import.meta.env.VITE_API_URL}/blog-comment`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					name: nameInput.value,
					comment: commentInput.value,
					blogID: $modalStore[0].meta.blogID
				})
			});

			if (response.ok) {
				toastStore.trigger({
					message: 'Comment submitted successfully',
					background: 'variant-filled-success',
					hideDismiss: true,
					timeout: 2000
				});
			} else {
				toastStore.trigger({
					message: 'Error submitting comment',
					background: 'variant-filled-error',
					hideDismiss: true,
					timeout: 2000
				});
			}
		} catch (error) {
			console.error('Error submitting comment:', error);
		}
		modalStore.close();
	}
</script>

<div class="p-4">
	<h2 class="text-xl font-bold">Add a Comment</h2>
	<form>
		<input id="name" type="text" placeholder="Your Name" class="input p-2 mb-2 w-full" />
		<textarea id="comment" placeholder="Your Comment" class="textarea p-2 mb-2 w-full"></textarea>
		<div class="flex gap-2 mt-4">
			<button type="submit" class="btn variant-glass-tertiary w-full" onclick={submitComment}
				>Submit</button
			>
			<button class="btn variant-glass-surface w-full" onclick={closeModal}>Cancel</button>
		</div>
	</form>
</div>
