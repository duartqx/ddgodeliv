import React, { useState } from "react";
import * as authService from "../services/auth/auth";
import { AuthContext } from "./AuthContext";
import * as Types from "./authTypes";

export default function AuthProvider({ children }) {
  const getAuth = () => {
    return JSON.parse(localStorage.getItem("auth") || "{}");
  };

  const [authData, setAuth] = useState(
    /** @type {Types.AuthData} */ (getAuth()),
  );
  const [user, setUser] = useState(/** @type {Types.User} */ ({}));

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

  return (
    <AuthContext.Provider value={{ user, authData, login, logout, register }}>
      {children}
    </AuthContext.Provider>
  );
}
