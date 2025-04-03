import React, { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { User, LoginCredentials, SignupCredentials, AuthContextType } from '../types/auth';

const BACKEND_URL = `${import.meta.env.VITE_BACKEND_HOST}:${import.meta.env.VITE_BACKEND_PORT}`;

const AuthContext = createContext<AuthContextType | null>(null);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      fetchUserData(token);
    }
  }, []);

  const fetchUserData = async (token: string) => {
    try {
      const response = await axios.get(`${BACKEND_URL}/me`, {
        headers: {
          Authorization: `Bearer ${token}`,
          AccessControlAllowOrigin: "*"
         }
      });
      setUser(response.data);
    } catch (error) {
      console.error('Failed to fetch user data:', error);
      localStorage.removeItem('token');
      setUser(null);
    }
  };

  const login = async (credentials: LoginCredentials) => {
    try {
      const response = await axios.post(`${BACKEND_URL}/login`, {
        username: credentials.username,
        password: credentials.password
      }, {
        headers: {
          AccessControlAllowOrigin: "*"
        }
      });
      const { token, user } = response.data;
      localStorage.setItem('token', token);
      setUser(user);
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    }
  };

  const signup = async (credentials: SignupCredentials) => {
    try {
      const response = await axios.post(`${BACKEND_URL}/signup`, {
        username: credentials.username,
        email: credentials.email,
        password: credentials.password
      }, {
        headers: {
          AccessControlAllowOrigin: "*"
        }
      });
      const { token, user } = response.data;
      localStorage.setItem('token', token);
      setUser(user);
    } catch (error) {
      console.error('Signup failed:', error);
      throw error;
    }
  };

  const logout = () => {
    localStorage.removeItem('token');
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ user, login, signup, logout, isAuthenticated: !!user }}>
      {children}
    </AuthContext.Provider>
  );
};
