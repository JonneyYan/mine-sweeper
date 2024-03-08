import React from "react";
import { render, act, screen } from "@testing-library/react";
import Counter from "../Counter";

describe("Counter", () => {
  beforeEach(() => {
    jest.useFakeTimers(); // set fake timer
  });

  afterEach(() => {
    jest.useRealTimers(); // reset real timer
  });
  it("render initial value", () => {
    render(<Counter />);
    expect(screen.getByText("0s")).toBeInTheDocument();
  });

  it("increases every second", () => {
    render(<Counter />);
    expect(screen.getByText("0s")).toBeInTheDocument();

    // wait 1s
    act(() => {
      jest.advanceTimersByTime(1000);
    });
    expect(screen.getByText("1s")).toBeInTheDocument();

    // wait 1s again
    act(() => {
      jest.advanceTimersByTime(1000);
    });
    expect(screen.getByText("2s")).toBeInTheDocument();
  });

  it("the timer stops when pause is true", () => {
    const { rerender } = render(<Counter pause={false} />);
    expect(screen.getByText("0s")).toBeInTheDocument();

    // wait 1s
    act(() => {
      jest.advanceTimersByTime(1000);
    });
    expect(screen.getByText("1s")).toBeInTheDocument();

    rerender(<Counter pause={true} />);

    // wait 1s again, should stop
    act(() => {
      jest.advanceTimersByTime(1000);
    });
    expect(screen.getByText("1s")).toBeInTheDocument(); // 计数不变
  });
});
