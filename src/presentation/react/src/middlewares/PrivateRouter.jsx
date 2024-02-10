import React from "react";
import { Navigate, useLocation } from "react-router-dom";
import useAuth from "../hooks/useAuth";

export default function PrivateRouter({ children }) {
  const { authData } = useAuth();
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
