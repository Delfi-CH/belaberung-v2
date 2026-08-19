import { isLoggedIn } from 'belaberung-client-libs/auth';
import { error, redirect } from '@sveltejs/kit';
import { createAPI, getBackendUrl, setBackendUrl } from 'belaberung-client-libs';

export const ssr = false;
export const prerender = false;

export async function load({ params }) {
	if (import.meta.env.DEV) { setBackendUrl("http://localhost:8081") }
	const api = createAPI(getBackendUrl())
	if (await isLoggedIn()) {
		try {
			const res = await api.get('/rooms/' + params.slug);
			return {
				post: res.data
			};
		} catch {
			error(500, 'server error');
		}
	} else {
		throw redirect(307, '/login');
	}
}
