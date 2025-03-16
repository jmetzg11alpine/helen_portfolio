<script>
	const data = [
		{
			title: "Gathering Painful Treasures: A Cargivers' Journey",
			image: '/treasures.jpg',
			content:
				'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Dui.',
			href: 'https://www.amazon.com/Gathering-Painful-Treasures-Caregivers-Journey-ebook/dp/B0D121JWQP/ref=sr_1_1?crid=28PK93M4G3ZSD&dib=eyJ2IjoiMSJ9.wiRJSq_1K67CUK4JX7xCRg.NHk0vsql_ypWmF_K4kyYIUUlLrvrrRu3jPbBh5xqvjM&dib_tag=se&keywords=a+caregivers+journey+helen+metzger&qid=1739235280&sprefix=a+caregivers+journey+helen+metzger%2Caps%2C71&sr=8-1',
			buttonText: 'Buy Now'
		},
		{
			title: 'Book Title 2',
			image: '/coming-soon.jpg',
			content:
				'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Dui.',
			href: '#newsletter',
			buttonText: 'Coming Soon'
		}
	];
</script>

<section id="books" class="mt-12 grid grid-cols-1 md:grid-cols-2 gap-8 items-stretch">
	{#each data as book}
		<div class="p-6 bg-surface-200 shadow-lg rounded-lg text-primary-500 flex flex-col">
			<h3 class="text-2xl font-bold mb-4">{book.title}</h3>
			<div class="aspect-[3/4] w-full max-w-[300px] mx-auto mb-4">
				<img
					src={book.image}
					alt="Book 1 Cover"
					class="w-full h-full object-cover rounded-lg shadow-md"
				/>
			</div>
			<p class="text-base-content flex-grow mb-4">
				{book.content}
			</p>
			<a href={book.href} class="btn text-white bg-secondary-500 self-center">
				{book.buttonText}
			</a>
		</div>
	{/each}
</section>
