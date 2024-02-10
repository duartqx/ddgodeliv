import { useState } from "react";

/**
 * @typedef {{
 *  token: ?string
 *  expiresAt: ?string
 *  status: ?string
 * }} AuthData
 */

export default function useAuth() {
  const getAuth = () => {
    return JSON.parse(localStorage.getItem("auth") || "{}");
  };

  const [authData, setAuth] = useState(/** @type {AuthData} */ (getAuth()));

  return { authData, setAuth };
}
