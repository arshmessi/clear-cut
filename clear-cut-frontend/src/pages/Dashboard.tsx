// src/pages/Dashboard.tsx
import React from "react";
import { Container, Typography, Box } from "@mui/material";

const Dashboard: React.FC = () => {
  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        Dashboard
      </Typography>
      <Box>
        {/* Include dashboard-specific components or content here */}
        <Typography variant="body1">
          Welcome to your dashboard! Here, you can manage your tasks, view
          analytics, and much more.
        </Typography>
      </Box>
    </Container>
  );
};

export default Dashboard;
