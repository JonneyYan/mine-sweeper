import React from "react";
import { render, fireEvent, screen } from "@testing-library/react";
import Level from "../Level";

describe("Level", () => {
  it("render level list", () => {
    render(<Level level={0} />);
    expect(screen.getByText("➤Primary")).toBeInTheDocument();
    expect(screen.getByText("Medium")).toBeInTheDocument();
    expect(screen.getByText("Expert")).toBeInTheDocument();
  });

  it("trigger onChange", () => {
    const handleChange = jest.fn();
    render(<Level level={0} onChange={handleChange} />);

    fireEvent.click(screen.getByText("Medium"));

    expect(handleChange).toHaveBeenCalledTimes(1);
    expect(handleChange).toHaveBeenCalledWith(1);
  });

  it("does not trigger onChange", () => {
    const handleChange = jest.fn();
    render(<Level level={0} disabled onChange={handleChange} />);

    // 点击 "Medium" 级别
    fireEvent.click(screen.getByText("Medium"));

    // 验证 onChange 回调是否未被调用
    expect(handleChange).not.toHaveBeenCalled();
  });
});
