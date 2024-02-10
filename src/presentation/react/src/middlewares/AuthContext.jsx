import React from "react";
import * as Types from "./authTypes";

export const AuthContext = React.createContext({
  user: /** {Types.User} */ {},
  authData: /** {Types.AuthData} */ {},
  /** @returns {Promise<Types.AuthData>} */
  login: async ({ email, password }) => {
    return /** @type {Promise<Types.AuthData>} */ ({});
  },

  logout: async () => {},

  /** @returns {Promise<boolean>} */
  register: async ({ name, email, password }) => {
    return false;
  },
});
