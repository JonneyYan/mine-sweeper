import React, { useEffect, useState } from "react";
import request from "../request";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { LEVEL_MAP } from "./Level";

dayjs.extend(relativeTime);

export default function History() {
  const [history, setHistory] = useState(null);
  useEffect(() => {
    fetchHistory();

    return () => {};
  }, []);

  async function fetchHistory() {
    const res = await request("/history");
    setHistory(res);
  }
  return (
    <>
      <button className="load-history" onClick={fetchHistory}>
        → Load History ←
      </button>
      <table className="history">
        <thead>
          <tr>
            <th>Username</th>
            <th>Steps</th>
            <th>Used Time</th>
            <th>Level</th>
            <th>Win</th>
          </tr>
        </thead>
        <tbody>
          {history?.map((item, i) => {
            return (
              <tr key={i}>
                <td>{item.username}</td>
                <td>{item.steps.length}</td>
                <td>{dayjs(item.endTime).diff(dayjs(item.startTime), "second")}s</td>
                <td>{LEVEL_MAP[item.difficulty]}</td>
                <td>{item.status === 2 ? "✓" : "×"}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </>
  );
}
