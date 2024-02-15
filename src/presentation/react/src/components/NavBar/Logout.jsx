import { useContext } from "react";
import { AuthContext } from "../../middlewares/AuthContext";
import { Link, useNavigate } from "react-router-dom";
import React from "react";

export default function Logout() {
  const navigate = useNavigate();
  const { logout } = useContext(AuthContext);

  const handleClick = async () => {
    await logout();
    navigate({ pathname: "/login" });
  };

  return (
    <Link to="#" className="text-danger" onClick={handleClick}>
      <i className="bi bi-box-arrow-right"></i>
    </Link>
  );
}
