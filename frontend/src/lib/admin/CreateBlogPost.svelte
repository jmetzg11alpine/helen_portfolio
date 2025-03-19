<script>
	let blogPost = $state({
		title: '',
		sub_title: '',
		content: ''
	});

	async function handleSubmit() {
		try {
			const url = `${import.meta.env.VITE_API_URL}/blog-new-entry`;
			console.log(url);
			const response = await fetch(url, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(blogPost)
			});

			if (response.ok) {
				blogPost = {
					title: '',
					sub_title: '',
					content: ''
				};
			}
		} catch (error) {
			console.error('Error creating blog post:', error);
		}
	}
</script>

<div class="card mx-auto p-6 bg-secondary-500 border border-secondary-200 my-4">
	<h2 class="h2 mb-4">Create New Blog Post</h2>

	<form onsubmit={preventDefault(handleSubmit)} class="space-y-4">
		<label class="label">
			<span>Title</span>
			<input
				class="input p-2"
				type="text"
				bind:value={blogPost.title}
				placeholder="Enter post title"
				required
			/>
		</label>

		<label class="label">
			<span>Subtitle</span>
			<input
				class="input p-2"
				type="text"
				bind:value={blogPost.sub_title}
				placeholder="Enter post subtitle"
			/>
		</label>

		<label class="label">
			<span>Content</span>
			<textarea
				class="textarea p-2"
				bind:value={blogPost.content}
				rows="6"
				placeholder="Write your blog post content here..."
				required
			></textarea>
		</label>

		<button type="submit" class="btn preset-filled-secondary-500 w-full"> Create Post </button>
	</form>
</div>
