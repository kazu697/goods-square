import React from 'react';
import './App.css';
import { BrowserRouter, Route } from 'react-router-dom';
import Home from './pages/Home';
import Header from './components/header/Header';
import ProductDetail from './pages/ProductDetail';

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Header>
        <Route exact path="/" component={Home} />
        <Route exact path="/product/:id" component={ProductDetail} />
      </Header>
    </BrowserRouter>
  );
};

export default App;
