import React from "react";
import { render, fireEvent, waitFor, screen, act } from "@testing-library/react";
import Minesweeper from "../Minesweeper";
import request from "../../request";

jest.mock("../../request", () => ({
  __esModule: true,
  default: jest.fn(),
}));
describe("Minesweeper", () => {
  beforeEach(() => {
    const mockConfData = {
      difficulties: {
        0: [8, 8, 10],
        1: [16, 16, 40],
        2: [16, 30, 99],
      },
    };
    request.mockResolvedValue(mockConfData);
  });

  it("render input and level selector", async () => {
    render(<Minesweeper />);
    await act(async () => {
      expect(request).toHaveBeenCalledTimes(1);
    });
    expect(screen.getByPlaceholderText("Please entry your name first...")).toBeInTheDocument();
    expect(screen.getByText("Primary")).toBeInTheDocument();
    expect(screen.getByText("Medium")).toBeInTheDocument();
    expect(screen.getByText("Expert")).toBeInTheDocument();
  });

  it("after entering the user name, the level selector is available", async () => {
    render(<Minesweeper />);
    await act(async () => {
      fireEvent.change(screen.getByPlaceholderText("Please entry your name first..."), { target: { value: "testUser" } });
    });
    expect(screen.getByText("Primary")).toBeEnabled();
    expect(screen.getByText("Medium")).toBeEnabled();
    expect(screen.getByText("Expert")).toBeEnabled();
  });

  it("trigger handleChangeLevel", async () => {
    render(<Minesweeper />);
    await act(async () => {
      expect(request).toHaveBeenCalledTimes(1);
    });
    fireEvent.change(screen.getByPlaceholderText("Please entry your name first..."), { target: { value: "testUser" } });
    fireEvent.click(screen.getByText("Medium"));

    expect(screen.getByText("➤Medium")).toBeInTheDocument();
  });

  it("trigger handleClick", async () => {
    render(<Minesweeper />);
    await act(async () => {
      fireEvent.change(screen.getByPlaceholderText("Please entry your name first..."), { target: { value: "testUser" } });
    });
    fireEvent.click(screen.getByText("Medium"));
    fireEvent.click(screen.getAllByTestId("mine-field")[0]);

    expect(screen.getAllByTestId("mine-field")).toHaveLength(256);

    expect(request).toHaveBeenCalled();
  });
});
