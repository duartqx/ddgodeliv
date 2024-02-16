import React, { useContext } from "react";
import { Navigate } from "react-router-dom";
import { AuthContext } from "./AuthContext";
import { Paths } from "../path";

export default function NotSignedInRouter({ children }) {
  const { isLoggedIn } = useContext(AuthContext);
  return isLoggedIn() ? <Navigate replace to={Paths.root} /> : children;
}
