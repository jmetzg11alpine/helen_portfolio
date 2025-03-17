<script>
	import { preventDefault } from 'svelte/legacy';

	import { fade } from 'svelte/transition';

	let name = $state(''),
		email = 'no-email',
		message = $state('');
	let success = $state(null);

	async function sendEmail() {
		const url = `${import.meta.env.VITE_API_URL}/contact`;
		const res = await fetch(url, {
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
	class="w-full max-w-6xl grid grid-cols-1 md:grid-cols-1 gap-8 mt-16 bg-surface-200 shadow-lg rounded-lg p-6 md:p-12 mx-auto"
	transition:fade
>
	<h2 class="text-3xl font-bold text-primary-500 text-center">Contact</h2>

	<form
		onsubmit={preventDefault(sendEmail)}
		class="space-y-6 max-w-2xl mx-auto w-full flex flex-col items-center"
	>
		<input
			bind:value={name}
			type="text"
			placeholder="Your Name"
			class="input input-bordered w-full px-4 py-3 rounded-lg"
			required
		/>
		<textarea
			bind:value={message}
			placeholder="Your message"
			class="textarea textarea-bordered w-full px-4 py-3 rounded-lg min-h-[150px]"
			required
		></textarea>
		<button
			type="submit"
			class="btn btn-primary bg-secondary-500 w-full md:w-1/2 py-3 text-lg font-semibold"
		>
			Contact
		</button>
	</form>
</section>

<!-- Success / Error Messages -->
{#if success === true}
	<p class="text-green-500 text-center mt-4">✅ Message sent successfully!</p>
{:else if success === false}
	<p class="text-red-500 text-center mt-4">❌ Failed to send message. Try again later.</p>
{/if}
