import { humanFileSize } from './utils';
import { Table, Col, Row, ProgressBar, Button } from 'react-bootstrap';
import './torrents.css';
import { Torrent } from './types/Torrent';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlay, faPause } from '@fortawesome/free-solid-svg-icons';
import usePostTorrentService from './services/usePostTorrentService';

type TorrentInfoProps = {
  torrent: Torrent;
};

const TorrentInfo: React.FC<TorrentInfoProps> = ({ torrent }) => {
  const { publishTorrent } = usePostTorrentService();
  let pauseTorrent = () => {
    torrent.isPaused = true;
    publishTorrent(torrent).then(() => console.log('updated'));
  };
  let playTorrent = () => {
    torrent.isPaused = false;
    publishTorrent(torrent).then(() => console.log('updated'));
  };
  return (
    <Row>
      <Col className="justify-content-start text-left">
        <ProgressBar
          now={Math.round(
            (100 * torrent.info.bytesCompleted) / torrent.info.length
          )}
          label={`${Math.round(
            (100 * torrent.info.bytesCompleted) / torrent.info.length
          )}%`}
          animated={!torrent.isPaused}
          variant={torrent.info.bytesMissing == 0 ? 'success' : 'info'}
          className="mt-3 mb-3"
        ></ProgressBar>

        <Row>
          <Col>
            {torrent.isPaused ? (
              <Button variant="success" onClick={playTorrent}>
                <FontAwesomeIcon icon={faPlay} />{' '}
              </Button>
            ) : (
              <Button variant="success" onClick={pauseTorrent}>
                <FontAwesomeIcon icon={faPause} />{' '}
              </Button>
            )}
          </Col>
        </Row>

        <Row>
          <Col xs={12} sm={6} md={5} lg={4}>
            <Table className="info-table">
              <thead></thead>
              <tbody>
                <tr>
                  <td>
                    <b>Down Speed:</b>
                  </td>
                  <td>TODO</td>
                </tr>
                <tr>
                  <td>
                    <b>Up Speed:</b>
                  </td>
                  <td>TODO</td>
                </tr>
                <tr>
                  <td>
                    <b>Downloaded:</b>
                  </td>
                  <td>{humanFileSize(torrent.info.bytesCompleted)}</td>
                </tr>
                <tr>
                  <td>
                    <b>Uploaded:</b>
                  </td>
                  <td>TODO</td>
                </tr>
              </tbody>
            </Table>
          </Col>
          <Col xs={12} sm={6} md={5} lg={4}>
            <Table className="info-table">
              <thead></thead>
              <tbody>
                <tr>
                  <td>
                    <b>Seeds:</b>
                  </td>
                  <td>TODO</td>
                </tr>
                <tr>
                  <td>
                    <b>Peers:</b>
                  </td>
                  <td>{torrent.info.peers}</td>
                </tr>
                <tr>
                  <td>
                    <b>Share Ratio:</b>
                  </td>
                  <td>TODO</td>
                </tr>
              </tbody>
            </Table>
          </Col>
        </Row>
      </Col>
    </Row>
  );
};

export default TorrentInfo;
