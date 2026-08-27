import { createContext, useState, useContext, useEffect } from 'react';
import { backend_url } from './BackendCall';

import { useNavigate } from 'react-router-dom';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(() => {
    const savedUser = localStorage.getItem('user')
    return savedUser ? JSON.parse(savedUser) : null
  });
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  const login = async (username, password) => {
    try {
      console.log('Attempting login...'); // Debug
      const response = await fetch(`${backend_url}/v1/login`, { // ← FIXED!
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
        credentials: 'include',
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Login failed');
      }

      const data = await response.json();
      console.log('Login response:', data); // Debug
      setUser(data.user);
      localStorage.setItem("user", JSON.stringify(data.user))
      navigate("/", { replace: true })
      return data.user;
    } catch (error) {
      console.error('Login error:', error);

      navigate("/", { replace: true })
      throw error;
    }
  };

  const logout = async () => {
    try {
      await fetch(`${backend_url}/v1/logout`, {
        method: 'POST',
        credentials: 'include',
      });
      setUser(null);
    } catch (error) {
      console.error('Logout error:', error);
    }
  };

  return (
    <AuthContext.Provider value={{ user, loading, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
