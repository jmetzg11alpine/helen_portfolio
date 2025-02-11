import nodemailer from 'nodemailer';
import { json } from '@sveltejs/kit';

export async function POST({ request }) {
	const { name, email, message } = await request.json();

	const transporter = nodemailer.createTransport({
		service: 'gmail',
		auth: {
			user: process.env.GMAIL_USER,
			pass: process.env.GMAIL_PASS
		}
	});

	let text;
	let subject;

	if (message === 'newsletter') {
		text = `add ${name} to news letter list`;
		subject = `News Letter - ${name}`;
	} else {
		text = message;
		subject = `Message from ${name}`;
	}

	const mailOptions = {
		from: email,
		to: process.env.GMAIL_USER,
		subject: subject,
		text: text
	};

	try {
		await transporter.sendMail(mailOptions);
		return json({ success: true, message: 'Email sent successfully!' });
	} catch (error) {
		console.error('Modemailer Error:', error);
		return json({ success: false, message: 'Failed to send email.' }, { status: 500 });
	}
}
