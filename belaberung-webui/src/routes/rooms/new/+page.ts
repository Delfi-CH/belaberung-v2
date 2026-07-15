import { isLoggedIn } from '$lib/api/auth';
import { redirect } from '@sveltejs/kit';

export async function load() {
	if (await isLoggedIn()) {
		return;
	} else {
		throw redirect(307, '/login');
	}
}
