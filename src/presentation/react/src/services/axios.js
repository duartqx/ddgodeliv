import axios from "axios";

const req = axios.create({
  baseURL: "http://127.0.0.1:8080",
});

req.interceptors.response.use(
  (res) => res,
  (error) => {
    if (error.response && error.response.status === 401) {
      sessionStorage.removeItem("auth");
    }
    return Promise.reject(error);
  },
);

export { req };
