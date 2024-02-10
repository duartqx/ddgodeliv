import axios from "axios";

export default function httpClient() {
  const client = axios.create({
    baseURL: "http://127.0.0.1:8080",
  });

  client.interceptors.request.use(
    (config) => {
      const authData = JSON.parse(localStorage.getItem("auth") || "{}");
      config.headers.Authorization = `Bearer ${authData?.token || ""}`;
      return config;
    },
    (error) => {
      return Promise.reject(error);
    },
  );

  client.interceptors.response.use(
    (res) => res,
    (error) => {
      if (error.response && error.response.status === 401) {
        localStorage.removeItem("auth");
      }
      return Promise.reject(error);
    },
  );
  return client;
}

export { httpClient };
