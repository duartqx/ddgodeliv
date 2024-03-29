import React, { useState } from "react";
import MainHeaderButton from "../../MainHeaderButton";

/**
 * @param {{
 *  driver: import("../../../services/driver/driver").Driver
 *  deleteHandler: (id: number) => void
 * }} props
 * */
export default function DriverMainHeader({ driver, deleteHandler }) {
  const [showConfig, setShowConfig] = useState(false);

  const onClickDeleteHandler = () => {
    if (
      window.confirm(`Are you sure you want to delete ${driver.user.name}?`)
    ) {
      deleteHandler(driver.id);
    }
  };

  return (
    <>
      <div className="my-4 d-flex" style={{ height: "8vh" }}>
        <div className="d-flex flex-column justify-content-center">
          <div className="fw-bold">{driver.user.name}</div>
          <div className="fw-light">{driver.user.email}</div>
          <div className="fw-light">LID: {driver.license_id}</div>
        </div>
        <div className="ms-auto d-flex">
          <div style={{ position: "relative" }}>
            <ul
              onMouseUp={() => setShowConfig(!showConfig)}
              onMouseLeave={() => setShowConfig(!showConfig)}
              className="dropdown-menu shadow"
              style={{
                position: "absolute",
                display: showConfig && "block",
                top: "100%",
              }}
            >
              <li role="button" className="dropdown-item">
                Edit
              </li>
              <li role="button" className="dropdown-item">
                Assign Vehicle
              </li>
              {driver.company.owner_id !== driver.user.id && (
                // The company owner MUST NOT be deleted
                <li
                  role="button"
                  className="dropdown-item"
                  onClick={onClickDeleteHandler}
                >
                  <span className="text-danger">Delete</span>
                </li>
              )}
            </ul>
          </div>
          <MainHeaderButton
            icon="bi-chat-left-text-fill"
            onClickHandler={() => alert("button chat")}
          />

          <MainHeaderButton
            icon="bi-telephone-fill"
            onClickHandler={() => alert("button phone")}
          />

          <MainHeaderButton
            icon="bi-three-dots-vertical"
            onClickHandler={() => setShowConfig(!showConfig)}
          />
        </div>
      </div>
    </>
  );
}
