import { api } from '$lib/api/core';
import axios from 'axios';

export async function isLoggedIn() {
	try {
		await api.get('/auth/status');
		return true;
	} catch {
		return false;
	}
}

export function getUsername(): string {
	const name = localStorage.getItem('username');
	if (!name) {
		return '';
	} else {
		return name;
	}
	return '';
}

export function getUserID(): string {
	const name = localStorage.getItem('userID');
	if (!name) {
		return '';
	} else {
		return name;
	}
	return '';
}

export type LoginResult = [success: boolean, errormsg: string];

export async function login(username: string, password: string): Promise<LoginResult> {
	let errorMessage = '';
	try {
		const res = await api.post('/auth/login', {
			username: username,
			password: password
		});
		localStorage.setItem('username', res.data.username);
		localStorage.setItem('userID', res.data.id);
		return [true, errorMessage];
	} catch (err) {
		if (axios.isAxiosError(err)) {
			if (err.status === 404) {
				errorMessage = `User ${username} doesnt exist!`;
			} else if (err.status === 418) {
				errorMessage = 'You are already logged in!';
				return [true, errorMessage];
			} else if (err.status === 400) {
				errorMessage = 'Malformed Request!';
			} else if (err.status === 500) {
				errorMessage = 'Internal server error!';
			} else if (err.status === 401) {
				errorMessage = `Incorrect Username or Password!`;
			} else {
				errorMessage = 'An unexpected error ocurred: ' + err;
			}
		}
	}
	return [false, errorMessage];
}

export async function register(username: string, password: string): Promise<LoginResult> {
	let errorMessage = '';
	try {
		const res = await api.post('/auth/register', {
			username: username,
			password: password
		});
		localStorage.setItem('username', res.data.username);
		localStorage.setItem('userID', res.data.id);
		return [true, errorMessage];
	} catch (err) {
		if (axios.isAxiosError(err)) {
			if (err.status === 404) {
				errorMessage = `User ${username} doesnt exist!`;
			} else if (err.status === 418) {
				errorMessage = 'You are already logged in!';
				return [true, errorMessage];
			} else if (err.status === 400) {
				errorMessage = 'Malformed Request!';
			} else if (err.status === 500) {
				errorMessage = 'Internal server error!';
			} else if (err.status === 409) {
				errorMessage = `Username ${username} is already in use!`;
			} else {
				errorMessage = 'An unexpected error ocurred: ' + err;
			}
		}
	}
	return [false, errorMessage];
}

export async function logout() {
	try {
		await api.get('/auth/logout');
		localStorage.removeItem('username');
		localStorage.removeItem('userID');
		return true;
	} catch {
		return false;
	}
}
