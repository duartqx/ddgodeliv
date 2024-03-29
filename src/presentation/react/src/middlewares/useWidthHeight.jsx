import { useState, useEffect } from "react";

export default function useWidthHeight() {
  const [windowWidth, setWindowWidth] = useState(window.innerWidth);
  const [windowHeight, setWindowHeight] = useState(window.innerHeight);

  const isSmallWindow = () => windowWidth <= 1200;

  useEffect(() => {
    const handleResize = () => {
      setWindowWidth(window.innerWidth);
      setWindowHeight(window.innerHeight);
    };
    window.addEventListener("resize", handleResize);
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return {
    windowWidth,
    setWindowWidth,
    windowHeight,
    setWindowHeight,
    isSmallWindow,
  };
}
