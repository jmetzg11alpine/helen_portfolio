<script>
	import { getModalStore, getToastStore } from '@skeletonlabs/skeleton';
	const modalStore = getModalStore();
	const toastStore = getToastStore();
	export let parent; // Suppresses the "unknown prop 'parent'" warning

	function closeModal() {
		modalStore.close();
	}

	function submitComment(event) {
		event.preventDefault();

		const info = $modalStore[0].meta;
		console.log(info);

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

		return;
	}
</script>

<div class="p-4">
	<h2 class="text-xl font-bold">Add a Comment</h2>
	<form>
		<input id="name" type="text" placeholder="Your Name" class="input mb-2 w-full" />
		<textarea id="comment" placeholder="Your Comment" class="textarea mb-2 w-full"></textarea>
		<div class="flex gap-2 mt-4">
			<button type="submit" class="btn variant-glass-tertiary w-full" on:click={submitComment}
				>Submit</button
			>
			<button class="btn variant-glass-surface w-full" on:click={closeModal}>Cancel</button>
		</div>
	</form>
</div>
