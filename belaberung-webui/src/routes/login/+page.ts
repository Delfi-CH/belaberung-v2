import { isLoggedIn } from 'belaberung-client-libs/auth';
import { redirect } from '@sveltejs/kit';
import { setBackendUrl } from 'belaberung-client-libs';

export async function load() {
	if (import.meta.env.DEV) { setBackendUrl("http://localhost:8081") }
	if (await isLoggedIn()) {
		throw redirect(307, '/');
	} else {
		return;
	}
}
