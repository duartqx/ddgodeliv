import React, { useContext } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { AuthContext } from "./AuthContext";

export default function PrivateRouter({ children }) {
  const { authData } = useContext(AuthContext);
  const location = useLocation();

  return authData?.status ? (
    children
  ) : (
    <Navigate
      replace
      to="/login"
      state={{ from: `${location.pathname}${location.search}` }}
    />
  );
}
