import React, { useEffect, useRef, useState } from "react";

function Counter({ pause }) {
  const [count, setCount] = useState(0);
  const intervalId = useRef();

  useEffect(() => {
    if (pause) {
      terminateTimer();
    }
  }, [pause]);

  useEffect(() => {
    startTimer();
    return () => clearInterval(intervalId.current);
  }, []);

  function startTimer() {
    intervalId.current = setInterval(() => {
      setCount((prevCount) => prevCount + 1);
    }, 1000);
  }
  function terminateTimer() {
    clearInterval(intervalId.current);
  }

  return (
    <div className="counter ">
      <span className="">{count}s</span>
    </div>
  );
}

export default Counter;
