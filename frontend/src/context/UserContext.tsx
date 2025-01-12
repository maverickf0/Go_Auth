import axios from 'axios';
import React, { createContext, useContext, useState } from 'react';

interface User {
  email: string;
}

interface UserContextType {
  user: User | null;
  login: (credentials: { email: string; password: string }) => Promise<void>;
  logout: () => void;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export function UserProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);

  const login = async (credentials: { email: string; password: string }) => {
    const response  = await axios.post('http://localhost:3000/auth/login', JSON.stringify(credentials));
    console.log(response)
    setUser({ email: credentials.email });
  };

  const logout = () => {
    setUser(null);
  };

  return (
    <UserContext.Provider value={{ user, login, logout }}>
      {children}
    </UserContext.Provider>
  );
}

export function useUser() {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return context;
}