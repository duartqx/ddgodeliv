import axios from "axios";

const baseUrl = "http://127.0.0.1:8080";

function getToken() {
  const authData = JSON.parse(localStorage.getItem("auth") || "{}");
  return authData?.token || "";
}

/** @returns {import("axios").AxiosInstance} */
export default function httpClient() {
  const client = axios.create({
    baseURL: baseUrl,
  });

  client.interceptors.request.use(
    (config) => {
      config.headers.Authorization = `Bearer ${getToken()}`;
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
