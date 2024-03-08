import React from "react";

function Scoreboard({ value }) {
  return (
    <div className="score-board">
      <div className="score-board-score">
        {value} {value > 1 ? "steps" : "step"}
      </div>
    </div>
  );
}

export default Scoreboard;
