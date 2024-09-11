// src/services/api.ts
const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080";

// src/services/api.ts
export const apiRequest = async (url: string, method: string, body?: any) => {
  const token = sessionStorage.getItem("token"); // This may not be necessary if using cookies

  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...(token && { Authorization: `Bearer ${token}` }), // Include token if necessary
  };

  const response = await fetch(`${API_URL}${url}`, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
    credentials: "include", // Important: This ensures cookies are included in the request
  });

  if (!response.ok) {
    throw new Error(`HTTP error! Status: ${response.status}`);
  }

  return response.json();
};
