import React from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch,
  useLocation,
} from 'react-router-dom';

import Header from './components/Header';
import About from './components/About';
import Feed from './components/Feed';
import Post from './components/Post';

function NoMatch() {
  let location = useLocation();

  return (
    <div class="wrap yue">
      <h3>
        No match for <code>{location.pathname}</code>
      </h3>
    </div>
  );
}

function App() {
  return (
    <Router>
      <Header />
      <Switch>
        <Route path="/posts/:permalink" component={Post} />
        <Route exact path="/about" exact component={About} />
        <Route exact path="/" component={Feed} />
        <Route path="*" component={NoMatch} />
      </Switch>
    </Router>
  );
}

export default App;
