import * as api from '$lib/api';

/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {
	const draw = await api.getDraw(fetch, params.uuid);
	return { draw };
}
