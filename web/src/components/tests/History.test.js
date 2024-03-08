import React, { render, fireEvent, waitFor, screen, act } from "@testing-library/react";
import History from "../History";
import request from "../../request.js";

jest.mock("../../request", () => ({
  __esModule: true,
  default: jest.fn(),
}));

describe("History", () => {
  beforeEach(() => {
    request.mockClear();
  });

  it("load history when component did mount", async () => {
    // mock data
    const mockHistory = [
      {
        username: "testUser",
        steps: [
          [0, 1],
          [0, 2],
          [0, 3],
        ],
        startTime: new Date("2022-01-01T00:00:00Z").toISOString(),
        endTime: new Date("2022-01-01T00:00:03Z").toISOString(),
        difficulty: 1,
        status: 2,
      },
    ];

    request.mockResolvedValue(mockHistory);

    render(<History />);
    
    await waitFor(() => {
      expect(request).toHaveBeenCalledTimes(1);
      expect(screen.getByText("testUser")).toBeInTheDocument();
      expect(screen.getByText("3")).toBeInTheDocument();
      expect(screen.getByText("3s")).toBeInTheDocument();
      expect(screen.getByText("✓")).toBeInTheDocument();
    });
  });

  it("load history when click button", async () => {
    render(<History />);

    fireEvent.click(screen.getByText("→ Load History ←"));

    await waitFor(() => {
      expect(request).toHaveBeenCalledTimes(2);
    });
  });
});
