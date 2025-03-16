<script>
	import '../app.postcss';
	import Header from '$lib/main_page/Header.svelte';
	import { initializeStores, Modal, Toast } from '@skeletonlabs/skeleton';
	import CommentModal from '$lib/blog/CommentModal.svelte';
	/**
	 * @typedef {Object} Props
	 * @property {import('svelte').Snippet} [children]
	 */

	/** @type {Props} */
	let { children } = $props();
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
	{@render children?.()}
</div>
