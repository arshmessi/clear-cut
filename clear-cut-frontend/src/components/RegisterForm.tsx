// src/components/RegisterForm.tsx
import React, { useState } from "react";
import { Button, TextField } from "@mui/material";
import { useAuth } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";

const RegisterForm: React.FC = () => {
  const { register } = useAuth();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      await register(name, email, password);
      navigate("/dashboard");
    } catch (error) {
      setError("Registration failed. Please try again.");
      console.error("Registration failed:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      {error && <div style={{ color: "red" }}>{error}</div>}
      <TextField
        label="Name"
        fullWidth
        value={name}
        onChange={(e) => setName(e.target.value)}
        margin="normal"
        required
      />
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
        Register
      </Button>
    </form>
  );
};

export default RegisterForm;
