import React, { useContext } from "react";
import { Navigate } from "react-router-dom";
import { AuthContext } from "./AuthContext";

export default function NotSignedInRouter({ children }) {
  const { authData } = useContext(AuthContext);
  return authData?.status ? <Navigate replace to="/dashboard" /> : children;
}
