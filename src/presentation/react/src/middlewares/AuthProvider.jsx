import React, { useState } from "react";
import * as authService from "../services/auth/auth";
import { AuthContext } from "./AuthContext";
import * as Types from "./authTypes";

export default function AuthProvider({ children }) {
  const [authData, setAuth] = useState(
    /** @type {Types.AuthData} */ (authService.getAuth()),
  );
  const [user, setUser] = useState(/** @type {Types.User} */ ({}));

  const getUser = async () => {
    if (!user?.id) {
      const u = await authService.getUser();
      setUser(u);
      return u;
    }
    return user;
  };

  /** @returns {Promise<Types.AuthData>} */
  const login = async ({ email, password }) => {
    const data = await authService.login({ email, password });
    setAuth(data);
    return data;
  };

  const logout = async () => {
    await authService.logout();
    setAuth(/** @type {Types.AuthData} */ ({}));
    setUser(/** @type {Types.User} */ ({}));
  };

  /** @returns {Promise<boolean>} */
  const register = async ({ name, email, password }) => {
    const signUpUser = await authService.register({ name, email, password });
    return signUpUser?.id ? true : false;
  };

  const isLoggedIn = () => Boolean(authData?.status);

  return (
    <AuthContext.Provider
      value={{ isLoggedIn, getUser, login, logout, register }}
    >
      {children}
    </AuthContext.Provider>
  );
}
