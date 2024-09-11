// src/context/AuthContext.tsx
import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  ReactNode,
} from "react";
import { apiRequest } from "../services/api";

type User = {
  id: number;
  name: string;
  email: string;
};

type AuthContextType = {
  user: User | null;
  login: (email: string, password: string) => Promise<void>;
  register: (name: string, email: string, password: string) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
};

type AuthProviderProps = {
  children: ReactNode;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const isAuthenticated = !!user;

  const login = async (email: string, password: string) => {
    try {
      const response = await apiRequest("/login", "POST", { email, password });
      sessionStorage.setItem("token", response.token);
      setUser(response.user);
    } catch (error) {
      console.error("Login failed:", error);
      alert("Login failed. Please check your credentials.");
    }
  };

  const register = async (name: string, email: string, password: string) => {
    try {
      const response = await apiRequest("/register", "POST", {
        name,
        email,
        password,
      });
      if (response.token) {
        sessionStorage.setItem("token", response.token);
        setUser(response.user);
      } else {
        console.log("Registration successful. Please log in.");
      }
    } catch (error) {
      console.error("Registration failed:", error);
    }
  };

  const logout = () => {
    setUser(null);
    sessionStorage.removeItem("token");
  };

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await apiRequest("/profile", "GET");
        setUser(response.user);
      } catch (error) {
        console.error("Failed to fetch user data:", error);
        sessionStorage.removeItem("token");
      }
    };

    fetchUserData();
  }, []);

  return (
    <AuthContext.Provider
      value={{ user, login, register, logout, isAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
};
