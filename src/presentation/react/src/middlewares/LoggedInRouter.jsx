import React from "react";
import { Navigate } from "react-router-dom";
import useAuth from "../hooks/useAuth";

export default function LoggedIn({ children }) {
  const [auth, _] = useAuth();
  return auth?.status ? <Navigate replace to="/dashboard" /> : children;
}
