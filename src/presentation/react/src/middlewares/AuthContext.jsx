import React from "react";
import * as Types from "./authTypes";

export const AuthContext = React.createContext({
  /** @returns {Promise<Types.User>} */
  getUser: async () => {
    return /** @type {Types.User} */ ({});
  },

  /** @returns {Promise<boolean>} */
  login: async ({ email, password }) => false,

  logout: async () => {},

  /** @returns {Promise<boolean>} */
  register: async ({ name, email, password }) => {
    return false;
  },

  /** @returns {boolean} */
  isLoggedIn: () => false,
});
