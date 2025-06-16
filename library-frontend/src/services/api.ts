// src/services/api.ts
import axios, { type AxiosResponse } from "axios";
import toast from "react-hot-toast";

const API_BASE_URL =
  import.meta.env.VITE_API_BASE_URL || "http://localhost:8080/api/v1";

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000,
});

// Request interceptor
api.interceptors.request.use(
  (config) => {
    console.log(
      `Making ${config.method?.toUpperCase()} request to ${config.url}`
    );
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  (response: AxiosResponse) => {
    return response;
  },
  (error) => {
    const message =
      error.response?.data?.error ||
      error.message ||
      "An unexpected error occurred";

    // Don't show toast for certain status codes
    if (error.response?.status !== 404) {
      toast.error(message);
    }

    return Promise.reject(error);
  }
);

export default api;
