import { useContext } from "react";
import { AuthContext } from "../middlewares/AuthContext";
import { useNavigate } from "react-router-dom";
import React from "react";

export default function Logout() {
  const navigate = useNavigate();
  const { logout } = useContext(AuthContext);

  const handleClick = async () => {
    await logout();
    navigate({ pathname: "/login" });
  };

  return (
    <button className="ms-auto btn btn-danger" onClick={handleClick}>
      Logout
    </button>
  );
}
