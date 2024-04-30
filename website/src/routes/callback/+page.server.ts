import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types.js';
import { mpuFetch, type Result } from '$lib/lib.js';

export interface Token {
	access_token: string;
	token_type: string;
	refresh_token: string;
	expiry: string;
}

export const load: PageServerLoad = async ({ url, cookies }) => {
	const code = url.searchParams.get('code');
	if (!code) {
		return { error: 'No code found in url' };
	}

	const res = (await mpuFetch('token?code=' + code, 'GET')) as Result;
	if (res.error) {
		console.log(res.error);
		return { error: res.error };
	}

	const token = res.data as Token;

	cookies.set('token', token.access_token, { path: '/' });

	redirect(302, '/');
};
