// src/pages/Login.tsx
import React from "react";
import { Container, Typography } from "@mui/material";
import LoginForm from "../components/LoginForm";

const Login: React.FC = () => {
  return (
    <Container maxWidth="sm">
      <Typography variant="h4" component="h1" gutterBottom>
        Login
      </Typography>
      <LoginForm />
    </Container>
  );
};

export default Login;
