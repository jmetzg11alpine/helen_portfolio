<script>
	import { authStore } from '$lib/admin/adminStore.js';
	let username = $state('');
	let password = $state('');
	let isLoading = $state(false);
	let errorMessage = $state('');
	let successMessage = $state('');

	async function handleSubmit(event) {
		event.preventDefault();
		isLoading = true;
		errorMessage = '';
		successMessage = '';

		try {
			const url = `${import.meta.env.VITE_API_URL}/login`;
			const resposne = await fetch(url, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					username: username,
					password: password
				})
			});

			if (!resposne.ok) {
				errorMessage = 'Invalid username or password';
			} else {
				successMessage = 'Login successful';
				const data = await resposne.json();

				authStore.set({
					token: data.token,
					role: data.role
				});
			}
			isLoading = false;
		} catch (error) {
			errorMessage = 'An error occurred';
			isLoading = false;
		}
	}
</script>

<div class="card mx-auto p-6 bg-surface-500 border border-surface-200 my-4">
	<form onsubmit={handleSubmit}>
		<label class="label">
			<span class="label-text text-lg">Username</span>
			<input
				class="input"
				type="text"
				placeholder="enter username"
				bind:value={username}
				required
			/>
		</label>
		<label class="label mt-4">
			<span class="label-text text-lg">Password</span>
			<input
				class="input"
				type="text"
				placeholder="enter password"
				bind:value={password}
				required
			/>
		</label>

		{#if errorMessage}
			<div class="test-error-500 my-2">{errorMessage}</div>
		{/if}

		{#if successMessage}
			<div class="test-success-500 my-2">{successMessage}</div>
		{/if}

		<button type="submit" class="btn preset-tonal-surface w-full mt-4" disabled={isLoading}>
			{#if isLoading}
				Loading...
			{:else}
				Login
			{/if}
		</button>
	</form>
</div>
