async function request(url, method = "GET", options = {}) {
  const baseURL = "";
  const resp = await fetch(`${baseURL}/api/v1${url}`, {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    body: options.body,
    ...options,
  });

  const res = await resp.json();

  if (res.code === 0) {
    return res.data;
  } else {
    console.log("🚀 ~ request ~ res.message:", res.message)
    // alert(res.message);
  }
}

export default request;
