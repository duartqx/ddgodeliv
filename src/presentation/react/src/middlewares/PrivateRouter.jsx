import React, { useContext } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { AuthContext } from "./AuthContext";

export default function PrivateRouter({ children }) {
  const { isLoggedIn } = useContext(AuthContext);
  const location = useLocation();

  return isLoggedIn() ? (
    children
  ) : (
    <Navigate
      replace
      to="/login"
      state={{ from: `${location.pathname}${location.search}` }}
    />
  );
}
