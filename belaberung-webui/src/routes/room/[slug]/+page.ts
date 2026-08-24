import { isLoggedIn } from 'belaberung-client-libs/auth';
import { error, redirect } from '@sveltejs/kit';
import { createAPI, getBackendUrl, setBackendUrl } from 'belaberung-client-libs';
import axios from 'axios';

export const ssr = false;
export const prerender = false;

export async function load({ params, url }) {
	if (import.meta.env.DEV) {
		setBackendUrl('http://localhost:8081');
	}
	const queryParams = url.searchParams;
	const addonParams = queryParams.get('password') ? `?password=${queryParams.get('password')}` : '';
	const api = createAPI(getBackendUrl());
	if (await isLoggedIn()) {
		try {
			const res = await api.get('/rooms/' + params.slug + addonParams);
			return {
				post: res.data,
				password: queryParams.get('password')
			};
		} catch (err) {
			if (err.status === 404 && queryParams.get('password')) {
				error(403, 'incorrect password');
			} else {
				switch (err.status) {
					case 400:
						error(400, 'bad request');
						break;
					case 401:
						error(401, 'unauthorised');
						break;
					case 403:
						error(403, 'forbidden');
						break;
					case 404:
						error(404, 'not found or missing password');
						break;
					default:
						error(500, 'server error');
				}
			}
		}
	} else {
		throw redirect(307, '/login');
	}
}
