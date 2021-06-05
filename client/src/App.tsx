import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import useTorrentsService from './services/useTorrentsService';

export default function BasicExample() {
  return (
    <Router>
      <div>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
        </ul>

        <hr />

        <Switch>
          <Route exact path="/">
            <Home />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

function Home() {
  return (
    <div>
      <h2>Home</h2>
      <Torrents />
    </div>
  );
}

const Torrents: React.FC<{}> = () => {
  const service = useTorrentsService();

  return (
    <div>
      {service.status === 'loading' && <div>Loading...</div>}
      {service.status === 'loaded' &&
        service.payload.results.map(torrent => (
          <div key={torrent.infoHash}>{torrent.infoHash}</div>
        ))}
      {service.status === 'error' && (
        <div>Error reaching backend: {service.error.toString()}</div>
      )}
    </div>
  );
};
