<script>
	import Header from '$lib/main_page/Header.svelte';
	import { ToastProvider, Modal } from '@skeletonlabs/skeleton-svelte';
	import CommentModal from '$lib/blog/CommentModal.svelte';
	import "../app.css";

	/**
	 * @typedef {Object} Props
	 * @property {import('svelte').Snippet} [children]
	 */
	/** @type {Props} */
	let { children } = $props();

	// 🚀 REMOVE initializeStores() – No longer needed in Skeleton 3

	const modalRegistry = {
		commentModal: { ref: CommentModal }
	};
</script>

<ToastProvider position="t" zIndex="z-[9999]" />
<Modal components={modalRegistry} />
<div
	class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-error-500 to-warning-500 text-center p-6"
>
	<Header />
	{@render children?.()}
</div>
