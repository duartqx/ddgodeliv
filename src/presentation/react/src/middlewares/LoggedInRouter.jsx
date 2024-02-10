import React from "react";
import { Navigate } from "react-router-dom";
import useAuth from "../hooks/useAuth";

export default function LoggedIn({ children }) {
  const { authData } = useAuth();
  return authData?.status ? <Navigate replace to="/dashboard" /> : children;
}
