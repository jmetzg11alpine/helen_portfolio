<script>
	import { fade } from 'svelte/transition';

	let name = '',
		email = '',
		message = 'newsletter';
	let success = null;

	async function sendEmail() {
		const res = await fetch('/api/contact', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ name, email, message })
		});

		const data = await res.json();
		success = data.success;
	}
</script>

<section
	id="newsletter"
	class="w-full max-w-6xl grid grid-cols-1 md:grid-cols-2 gap-12 mt-16 bg-surface-200 shadow-lg rounded-lg p-6 md:p-12"
	transition:fade
>
	<div class="flex flex-col justify-center text-center md:text-left">
		<h2 class="text-3xl font-bold text-primary-500">Stay Connected</h2>
		<p class="mt-4 text-base-content">
			Subscribe to my newsletter for the latest updates and extra content.
		</p>
	</div>

	<form on:submit|preventDefault={sendEmail} class="space-y-4 max-w-lg mx-auto w-full">
		<input
			bind:value={name}
			type="text"
			placeholder="Your Name"
			class="input input-bordered w-full px-4 py-3 rounded-lg"
			required
		/>
		<input
			bind:value={email}
			type="email"
			placeholder="Your Email"
			class="input input-bordered w-full px-4 py-3 rounded-lg"
			required
		/>
		<button type="submit" class="btn btn-primary bg-secondary-500 w-full py-3 text-lg font-semibold"
			>Subscribe</button
		>
	</form>
</section>

<!-- Success / Error Messages -->
{#if success === true}
	<p class="text-green-500 text-center mt-4">✅ Email sent successfully!</p>
{:else if success === false}
	<p class="text-red-500 text-center mt-4">❌ Failed to send email. Try again later.</p>
{/if}
