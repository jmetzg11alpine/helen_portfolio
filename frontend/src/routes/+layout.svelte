<script>
	import '../app.postcss';
	import Header from '$lib/main/Header.svelte';
	import { initializeStores, Modal, Toast } from '@skeletonlabs/skeleton';
	import CommentModal from '$lib/blog/CommentModal.svelte';
	initializeStores();

	const modalRegistry = {
		commentModal: { ref: CommentModal }
	};
</script>

<Toast position="t" zIndex="z-[9999]" />
<Modal components={modalRegistry} />
<div
	class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br variant-gradient-error-warning text-center p-6"
>
	<Header />
	<slot />
</div>
