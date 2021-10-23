export async function makeFetchHappen(searchValue) {
  try {
    const res = await fetch("http://127.0.0.1:3001/user-search", {
      method: "POST",
      body: JSON.stringify({ search: searchValue }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const json = await res.json();
    return json;
  } catch (err) {
    console.error(err);
  }
}
