import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import Header from './components/Header';
import About from './components/About';
import Feed from './components/Feed';
import Post from './components/Post';

function App() {
  return (
    <Router>
      <Header />
      <Route path="/posts/:permalink" component={Post} />
      <Route path="/about" exact component={About} />
      <Route path="/" exact component={Feed} />
    </Router>
  );
}

export default App;
