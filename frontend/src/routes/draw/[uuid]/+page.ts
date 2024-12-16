import * as api from '$lib/api';

export const prerender = false;

type RouteParams = {
  uuid: string;
};

/** @type {import('./$types').PageLoad} */
export async function load({ params }: { params: RouteParams }) {
	const draw = await api.getDraw(params.uuid);
	return { draw };
}
