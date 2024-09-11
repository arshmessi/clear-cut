// src/pages/Register.tsx
import React from "react";
import { Container, Typography } from "@mui/material";
import RegisterForm from "../components/RegisterForm";

const Register: React.FC = () => {
  return (
    <Container maxWidth="sm">
      <Typography variant="h4" component="h1" gutterBottom>
        Register
      </Typography>
      <RegisterForm />
    </Container>
  );
};

export default Register;
