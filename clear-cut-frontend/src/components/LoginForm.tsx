// src/components/LoginForm.tsx
import React, { useState } from "react";
import { Button, TextField } from "@mui/material";
import { useAuth } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";

const LoginForm: React.FC = () => {
  const { login } = useAuth();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      await login(email, password);
      navigate("/dashboard"); // Redirect to dashboard after successful login
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        label="Email"
        type="email"
        fullWidth
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        margin="normal"
        required
      />
      <TextField
        label="Password"
        type="password"
        fullWidth
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        margin="normal"
        required
      />
      <Button type="submit" variant="contained" color="primary" fullWidth>
        Login
      </Button>
    </form>
  );
};

export default LoginForm;
