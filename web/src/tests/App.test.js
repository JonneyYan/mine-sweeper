import { render, screen } from '@testing-library/react';
import App from '../App';

test('renders game name', () => {
  render(<App />);
  const linkElement = screen.getByText(/Minesweeper Online/i);
  expect(linkElement).toBeInTheDocument();
});
