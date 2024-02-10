import { useState } from "react";

/**
 * @typedef {{
 *  token: ?string
 *  expiresAt: ?string
 *  status: ?string
 *  user: ?{
 *    id: number
 *    email: string
 *    name: string
 *    driver: {
 *      driver_id: number
 *      company_id: number
 *    }
 *  }
 * }} AuthData
 */

export default function useAuth() {
  const getAuth = () => {
    return JSON.parse(localStorage.getItem("auth") || "{}");
  };

  const [auth, setAuth] = useState(/** @type {AuthData} */ (getAuth()));

  return [auth, setAuth];
}
