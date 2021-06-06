import { useEffect, useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { Row, Col } from 'react-bootstrap';
import TorrentsList from './TorrentsList';
import TorrentInfo from './TorrentInfo';
import { Torrent } from './types/Torrent';
import useTorrentsService, { Torrents } from './services/useTorrentsService';
import { Service } from './types/Service';

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
  const [selectedInfoHash, setSelectedInfoHash] = useState<string>('');
  const {
    service,
    refreshTorrents,
  }: { service: Service<Torrents>; refreshTorrents: () => void } =
    useTorrentsService();

  useEffect(() => {
    const interval = setInterval(() => refreshTorrents(), 1000);
    return () => {
      clearInterval(interval);
    };
  }, []);

  const selectedTorrent =
    service.status == 'loaded' &&
    service.payload.results.find(
      (torrent: Torrent) => torrent.infoHash === selectedInfoHash
    );

  return (
    <Row className="justify-content-center h-100">
      <Col xs={2} className="border"></Col>
      <Col xs={10}>
        <div className="h-100 d-flex flex-column">
          <Row className="justify-content-center flex-grow-1 border">
            {service.status === 'loading' && (
              <tr>
                <td>Loading...</td>
              </tr>
            )}

            {service.status === 'loaded' && (
              <TorrentsList
                torrents={service.payload.results}
                selectedInfoHash={selectedInfoHash}
                setSelectedInfoHash={setSelectedInfoHash}
              />
            )}
          </Row>

          {service.status === 'loaded' && selectedTorrent && (
            <TorrentInfo torrent={selectedTorrent} />
          )}

          {service.status === 'error' && (
            <div>Error reaching backend: {service.error.toString()}</div>
          )}
        </div>
      </Col>
    </Row>
  );
}
