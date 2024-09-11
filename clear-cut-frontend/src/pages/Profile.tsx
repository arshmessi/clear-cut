// src/pages/Profile.tsx
import React from "react";
import { Container, Typography } from "@mui/material";
import ProfileDetails from "../components/ProfileDetails";

const Profile: React.FC = () => {
  return (
    <Container maxWidth="sm">
      <Typography variant="h4" component="h1" gutterBottom>
        User Profile
      </Typography>
      <ProfileDetails />
    </Container>
  );
};

export default Profile;
