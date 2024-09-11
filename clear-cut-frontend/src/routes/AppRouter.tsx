// src/AppRouter.tsx
import React from "react";
import { Route, Routes, Navigate } from "react-router-dom";
import Dashboard from "../pages/Dashboard";
import Login from "../pages/Login";
import Register from "../pages/Register";
import Profile from "../pages/Profile";
import { useAuth } from "../context/AuthContext";

const PrivateRoute: React.FC<{ element: React.ReactElement }> = ({
  element,
}) => {
  const { isAuthenticated } = useAuth();
  return isAuthenticated ? element : <Navigate to="/login" />;
};

const AppRouter: React.FC = () => {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/profile" element={<PrivateRoute element={<Profile />} />} />
      <Route
        path="/dashboard"
        element={<PrivateRoute element={<Dashboard />} />}
      />
      <Route path="/" element={<Navigate to="/dashboard" />} />{" "}
      {/* Redirect to dashboard by default */}
    </Routes>
  );
};

export default AppRouter;
