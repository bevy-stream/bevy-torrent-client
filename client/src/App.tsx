import { useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { Row, Col } from 'react-bootstrap';
import Torrents from './Torrents';
import TorrentInfo from './TorrentInfo';
import { Torrent } from './types/Torrent';

export default function Site() {
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <Home />
        </Route>
      </Switch>
    </Router>
  );
}

function Home() {
  const [selected, setSelected] = useState<Torrent | undefined>(undefined);

  return (
    <Row className="justify-content-center h-100">
      <Col xs={2} className="border"></Col>
      <Col xs={10}>
        <div className="h-100 d-flex flex-column">
          <Row className="justify-content-center flex-grow-1 border">
            <Torrents selected={selected} setSelected={setSelected} />
          </Row>
          {selected && <TorrentInfo torrent={selected} />}
        </div>
      </Col>
    </Row>
  );
}
