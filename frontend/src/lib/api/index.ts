export type Draw = {
  uuid: string;
  config?: {
    chained?: boolean;
  };
};

export type Member = {
  uuid?: string;
  name: string;
};

export async function createDraw() {
  // api post request to localhost:8080/api/v1/draw
  return await fetch('http://localhost:8080/api/v1/draw', {
    method: 'POST',
  })
  .then(response => response.json())
  .then(data => {
    return data;
  })
  .catch((error) => {
    console.error('Error:', error);
  });
}

export async function getDraw(uuid: string) {
  return await fetch('http://localhost:8080/api/v1/draw/' + uuid, {
    method: 'GET'
  })
  .then(response => response.json())
  .then(data => {
    return data;
  })
  .catch((error) => {
    console.error('Error:', error);
  });
}

export async function patchDrawConfig(uuid: string, draw: Draw) {
  return await fetch('http://localhost:8080/api/v1/draw/' + uuid, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(draw)
  })
  .then(response => response.json())
  .then(data => {
    return data;
  })
  .catch((error) => {
    console.error('Error:', error);
  });
}

export async function createDrawMember(uuid: string, member: Member) {
  return await fetch('http://localhost:8080/api/v1/draw/' + uuid + '/member', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(member)
  })
  .then(response => response.json())
  .then(data => {
    return data;
  })
  .catch((error) => {
    console.error('Error:', error);
  });
}

export async function deleteMember(uuid: string) {
  return await fetch('http://localhost:8080/api/v1/member/' + uuid, {
    method: 'DELETE'
  })
  .then(response => response.json())
  .catch((error) => {
    console.error('Error:', error);
  });
}
