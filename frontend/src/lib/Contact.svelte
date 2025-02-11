<script>
	import { fade } from 'svelte/transition';

	let name = '',
		email = 'no-email',
		message = '';
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
	id="contact"
	class="w-full max-w-6xl flex flex-col align-middle gap-8 mt-16 bg-surface-200 shadow-lg rounded-lg px-12 py-12"
	transition:fade
>
	<h2 class="text-3xl font-bold text-primary-500 text-center">Contact</h2>

	<div class="w-full max-w-6xl flex flex-col md:flex-row gap-12 justify-between">
		<textarea
			bind:value={message}
			placeholder="Your message"
			class="textarea textarea-bordered w-full px-4 py-3 rounded-lg min-h-[150px] flex-1"
			required
		></textarea>
		<form
			on:submit|preventDefault={sendEmail}
			class="space-y-4 max-w-lg mx-auto w-full flex flex-col justify-between h-full flex-1"
		>
			<input
				bind:value={name}
				type="text"
				placeholder="Your Name"
				class="input input-bordered w-full px-4 py-3 rounded-lg"
				required
			/>
			<button
				type="submit"
				class="btn btn-primary bg-secondary-500 w-full py-3 text-lg font-semibold"
			>
				Contact
			</button>
		</form>
	</div>
</section>

<!-- Success / Error Messages -->
{#if success === true}
	<p class="text-green-500 text-center mt-4">✅ Message sent successfully!</p>
{:else if success === false}
	<p class="text-red-500 text-center mt-4">❌ Failed to send message. Try again later.</p>
{/if}
