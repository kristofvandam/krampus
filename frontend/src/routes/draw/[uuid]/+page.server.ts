import * as api from '$lib/api';

export const actions = {
	create_member: async ({ request }) => {
		const data = await request.formData();
    const uuid = data.get('uuid');
    const name = data.get('name');
    api.createDrawMember(uuid, {
      name: name,
    });
	},
  delete_member: async ({ request }) => {
    const data = await request.formData();
    const uuid = data.get('uuid');
    api.deleteMember(uuid);
  }
};
