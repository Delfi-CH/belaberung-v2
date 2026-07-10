import { isLoggedIn } from '$lib/api/auth';
import { redirect } from '@sveltejs/kit';

export async function load() {
	if (await isLoggedIn()) {
		throw redirect(307, '/');
	} else {
		return;
	}
}
