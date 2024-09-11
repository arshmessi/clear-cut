// src/components/ProfileDetails.tsx
import React from "react";
import { Box, Typography, Button } from "@mui/material";
import { useAuth } from "../context/AuthContext";

const ProfileDetails: React.FC = () => {
  const { user, logout } = useAuth();

  return (
    <Box>
      <Typography variant="h5">Profile Information</Typography>
      {user && (
        <>
          <Typography variant="body1">Name: {user.name}</Typography>
          <Typography variant="body1">Email: {user.email}</Typography>
        </>
      )}
      <Button variant="contained" color="secondary" onClick={logout}>
        Logout
      </Button>
    </Box>
  );
};

export default ProfileDetails;
