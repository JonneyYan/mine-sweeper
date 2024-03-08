import React, { memo, useCallback, useEffect, useMemo, useState } from "react";
import times from "lodash/times";

import Scoreboard from "./Scoreboard";
import Counter from "./Counter";
import Level from "./Level";
import request from "../request";

const MinesweeperStatus = {
  Live: 0,
  GameOver: 1,
  Win: 2,
};

function Minesweeper() {
  const [size, setSize] = useState({ rows: 10, cols: 10 });
  const [username, setUsername] = useState(undefined);
  const [difficulty, setDifficulty] = useState([]);
  const [level, setLevel] = useState(null);
  const [game, setGame] = useState(null);
  const [loading, setLoading] = useState(null);

  useEffect(() => {
    async function fetchConf() {
      const data = await request("/conf");
      setDifficulty(data.difficulties);
    }
    fetchConf();
  }, []);

  const handleChangeLevel = useCallback(
    (level) => {
      if (difficulty.length === 0) {
        return;
      }
      setGame(null);
      setLevel(level);
      setSize({
        rows: difficulty[level][0],
        cols: difficulty[level][1],
      });
    },
    [difficulty]
  );

  const isLiving = useMemo(() => {
    if (level == null || !username || (game && game.status !== MinesweeperStatus.Live)) {
      return false;
    } else {
      return true;
    }
  }, [level, username, game]);

  const handleClick = useCallback(
    (i, j) => {
      return async () => {
        try {
          if (!!loading) return;
          setLoading([i, j]);
          if (!game) {
            if (level == null || !username) {
              return;
            }
            const data = await request("/start", "POST", {
              body: JSON.stringify({
                username,
                difficulty: level,
                initialPosition: [i, j],
              }),
            });
            setGame(data);
          } else {
            if (!isLiving) {
              return;
            }
            const data = await request("/game/clear", "POST", {
              body: JSON.stringify({
                x: i,
                y: j,
              }),
            });
            setGame(data);
          }
        } finally {
          setLoading(null);
        }
      };
    },
    [loading, isLiving, level, username, game]
  );
  const handleRightClick = useCallback(
    (i, j) => {
      return async (e) => {
        try {
          if (!!loading) return;
          setLoading([i, j]);
          e.preventDefault();
          if (!isLiving) {
            return;
          }
          const data = await request("/game/flagging", "POST", {
            body: JSON.stringify({
              x: i,
              y: j,
            }),
          });
          setGame(data);
        } finally {
          setLoading(null);
        }
      };
    },
    [isLiving, loading]
  );
  const handleDoubleClick = useCallback(
    (i, j) => {
      return async (e) => {
        try {
          if (!!loading) return;
          setLoading([i, j]);
          e.preventDefault();
          if (!isLiving) {
            return;
          }
          const data = await request("/game/clearAdjacent", "POST", {
            body: JSON.stringify({
              x: i,
              y: j,
            }),
          });
          setGame(data);
        } finally {
          setLoading(null);
        }
      };
    },
    [isLiving, loading]
  );

  return (
    <div className="minesweeper">
      <div className="mine-username">
        <input className="username" type="text" placeholder="Please entry your name first..." value={username} onChange={(e) => setUsername(e.target.value)} />
      </div>
      <div className="mine-board">
        {game && <Scoreboard value={game.steps.length} />}
        <Level level={level} disabled={!username} onChange={handleChangeLevel} />
        {game && <Counter pause={game.status !== MinesweeperStatus.Live} />}
      </div>
      <div className="mine-field">
        {times(size.rows, (i) => (
          <div key={i} className="mine-field-rows" style={{ gridTemplate: `repeat(1, 1fr) / repeat(${size.cols}, 1fr)` }}>
            {times(size.cols, (j) => {
              const cell = game?.grid[i][j];
              let v = cell ? (cell.flagging ? "🚩" : cell?.value) : " ";

              if (loading !== null && i === loading[0] && j === loading[1]) {
                v = "🧭";
              }
              return (
                <span
                  data-testid="mine-field"
                  key={j}
                  onClick={handleClick(i, j)}
                  onContextMenu={handleRightClick(i, j)}
                  onDoubleClick={handleDoubleClick(i, j)}
                  className={`mine-filed-item mine-filed-item-${v}`}
                >
                  {v !== "0" && v !== "❔" ? v : "  "}
                </span>
              );
            })}
          </div>
        ))}
      </div>
    </div>
  );
}

export default memo(Minesweeper);
