import React, { useCallback } from "react";

export const LEVEL_MAP = ["Primary", "Medium", "Expert"];
function Level({ level, disabled, onChange }) {
  const handleClick = useCallback(
    (index) => {
      return () => {
        if (disabled) return;
        onChange(index);
      };
    },
    [disabled, onChange]
  );

  return (
    <div className={`level ${disabled ? "disabled" : ""}`}>
      <h1>Select Level To (Re)Start</h1>
      <div className="level-select">
        {LEVEL_MAP.map((item, index) => {
          if (index === level) {
            return (
              <span key={index} className="active" onClick={handleClick(index)}>
                ➤{item}
              </span>
            );
          }
          return (
            <span key={index} onClick={handleClick(index)}>
              {item}
            </span>
          );
        })}
      </div>
    </div>
  );
}

export default Level;
