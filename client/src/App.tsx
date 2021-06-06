import { useState } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import { Row, Col } from 'react-bootstrap';
import Torrents from './Torrents';

export default function Site() {
  return (
    <Router>
      <div className="h-100">
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
  const [selected, setSelected] = useState(null);

  return (
    <Row className="h-100">
      <Col xs={2} className="border"></Col>
      <Col xs={10} className="">
        <div className="h-100">
          <Row className="h-75">
            <Torrents selected={selected} setSelected={setSelected} />
          </Row>
          <Row className="border-top h-25">
            <Col xs={12} className="">
            </Col>
          </Row>
        </div>
      </Col>
    </Row>
  );
}

