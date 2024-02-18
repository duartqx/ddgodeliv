import React from "react";
import useWidthHeight from "../../middlewares/useWidthHeight";

export default function CardFormCreateButton({
  label,
  height,
  onClickHandler,
}) {
  const {windowHeight, windowWidth } = useWidthHeight()

  const getBottom = () => {
    console.log(windowHeight, windowWidth)
    if (windowWidth > 800) {
      return 0
    }
    return "-1rem"
  }

  return (
      <div
        className="p-3 d-flex justify-content-center"
        style={{
          position: "absolute",
          background: "linear-gradient(transparent 0%, white 50%)",
          width: "22rem",
          height: `${height}`,
          bottom: getBottom()
        }}
      >
        <button
          className="btn align-self-end text-white p-3 shadow"
          style={{
            backgroundColor: "#000",
            width: "100%",
          }}
          onClick={onClickHandler}
        >
          {label}
        </button>
      </div>
  );
}
