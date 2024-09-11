// src/services/auth.ts
export async function login(email: string, password: string) {
  const response = await fetch(`${process.env.REACT_APP_API_URL}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
    credentials: "include", // Include cookies in the request
  });

  if (!response.ok) {
    throw new Error("Failed to log in");
  }

  // No need to handle token explicitly if it's managed by cookies
}
