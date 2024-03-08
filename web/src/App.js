import "./App.css";
import Minesweeper from "./components/Minesweeper";
import History from "./components/History";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Minesweeper Online</h1>
      </header>
      <Minesweeper />
      <History />
    </div>
  );
}

export default App;
