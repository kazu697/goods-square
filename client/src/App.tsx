import "./App.css";
import { BrowserRouter, Route } from "react-router-dom";
import Home from "./pages/Home";
import Header from "./components/header/Header";

function App() {
  return (
    <BrowserRouter>
      <Header>
        <Route exact path="/" component={Home} />
      </Header>
    </BrowserRouter>
  );
}

export default App;
