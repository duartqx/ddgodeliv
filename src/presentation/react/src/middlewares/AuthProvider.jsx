import React, { useState } from "react";
import * as authService from "../services/auth/auth";
import { AuthContext } from "./AuthContext";
import * as Types from "./authTypes";

export default function AuthProvider({ children }) {
  const emptyUser = /** @type {Types.User} */ ({});

  const [user, setUser] = useState(emptyUser);

  /** @returns {Promise<boolean>} */
  const login = async ({ email, password }) => {
    const data = await authService.login({ email, password });
    return Boolean(data.status);
  };

  const logout = async () => {
    await authService.logout();
    setUser(emptyUser);
  };

  /** @returns {Promise<boolean>} */
  const register = async ({ name, email, password }) => {
    const signUpUser = await authService.register({ name, email, password });
    return signUpUser?.id ? true : false;
  };

  const getUser = async () => {
    const exp = authService.getAuth()?.expiresAt
    if (!exp) {
      return emptyUser
    }

    if ((new Date(exp)) < (new Date())) {
      logout();
      return emptyUser;
    }

    if (!user?.id) {
      const authUser = await authService.getUser();
      if (!authUser?.id) {
        logout();
      } else {
        setUser(authUser);
      }
    }
    return user;
  };

  const isLoggedIn = () => {
    const exp = authService.getAuth()?.expiresAt;
    if (!exp) {
      return false
    }
    return Boolean((new Date(exp)) > (new Date()))
  };

  return (
    <AuthContext.Provider
      value={{ isLoggedIn, getUser, login, logout, register }}
    >
      {children}
    </AuthContext.Provider>
  );
}
